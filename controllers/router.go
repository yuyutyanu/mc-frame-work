package controllers

import (
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

type controllerInfo struct {
	regex *regexp.Regexp
	params map[int]string
	controllerType reflect.Type
}
type ControllerRegsiter struct{
	routes []*controllerInfo
	Application *App
}

type App struct {}
type Context struct {
	ResponseWriter http.ResponseWriter
	Request *http.Request
	Params map[string]string
}

func (p *ControllerRegsiter) Add(pettern string, c ControllerInterface){
	parts := strings.Split(pettern,"/")
	params:= make(map[int]string)

	j:=0
	for i, part := range parts{
		// '/part/:param'
		if strings.HasPrefix(part, ":"){
			expr := "([^/]+)"
			// '/user/:id([0-9]+)' などの正規表現対応
			if index := strings.Index(part, "("); index != -1 {
				expr = part[index:]
				part = part[:index]
			}

			params[j] = part
			parts[i] = expr
			j++
		}
	}
	newPattern := strings.Join(parts, "/")
	regex, regexErr := regexp.Compile(newPattern)

	if regexErr != nil {
		panic(regexErr)
		return
	}

	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	route := &controllerInfo{}
	route.regex = regex
	route.params = params
	route.controllerType = t

	p.routes = append(p.routes, route)
}

func (p *ControllerRegsiter) ServeHTTP(w http.ResponseWriter, r *http.Request){

	//Todo :------------------------
	//defer func(){
	//	if err := recover(); err != nil{
	//		if !RecoverPanic {
	//			panic(err)
	//		}else {
	//			Critical("Handler crashed with error", err)
	//			for i := 1; ;i +=1  {
	//				_, file, line,ok := runtime.Caller(i)
	//				if !ok {
	//					break;
	//				}
	//				Critical(file, line)
	//			}
	//		}
	//	}
	//}()
	// -----------------------------:Todo

	var StaticDir map[string]string
	// static routing
	var started bool
	requestPath := r.URL.Path
	for prefix, staticDir := range StaticDir{
		//完全一致に変えたほうがいい　/imghogehogeで通ってしまう
		if strings.HasPrefix(requestPath, prefix){
			file := staticDir + requestPath[len(prefix):]
			http.ServeFile(w, r, file)
			started = true
			return
		}
	}

	for _, route := range p.routes{
		if !route.regex.MatchString(requestPath){
			continue
		}

		// /:id..etc params join query string
		matches := route.regex.FindStringSubmatch(requestPath)
		if len(matches[0]) != len(requestPath){
			continue
		}

		params := make(map[string]string)
		if len(route.params) > 0{
			values := r.URL.Query()
			for i, match := range matches[1:]{
				values.Add(route.params[i], match)
				params[route.params[i]] = match
			}
			r.URL.RawQuery = url.Values(values).Encode() + "&" + r.URL.RawQuery
		}


		vc := reflect.New(route.controllerType) // vc is reflect.Value of Controller
		init := vc.MethodByName("Init")
		in := make([]reflect.Value, 2)
		ct := &Context{ResponseWriter: w, Request: r, Params: params}
		in[0] = reflect.ValueOf(ct)
		in[1] = reflect.ValueOf(route.controllerType.Name())
		init.Call(in)
		in = make([]reflect.Value, 0)
		method := vc.MethodByName("Prepare")
		method.Call(in)

		// if　はださいので　route.getみたいに叩ける用変える
		if r.Method == "GET"{
			method = vc.MethodByName("Get")
			method.Call(in)
		}else if r.Method == "POST" {
			method = vc.MethodByName("Post")
			method.Call(in)
		} else if r.Method == "HEAD" {
			method = vc.MethodByName("Head")
			method.Call(in)
		} else if r.Method == "DELETE" {
			method = vc.MethodByName("Delete")
			method.Call(in)
		} else if r.Method == "PUT" {
			method = vc.MethodByName("Put")
			method.Call(in)
		} else if r.Method == "PATCH" {
			method = vc.MethodByName("Patch")
			method.Call(in)
		} else if r.Method == "OPTIONS" {
			method = vc.MethodByName("Options")
			method.Call(in)
		}

		if AutoRender := false; AutoRender {
			method = vc.MethodByName("Render")
			method.Call(in)
		}
		method = vc.MethodByName("Finish")
		method.Call(in)
		started = true
		break
	}

	if started == false{
		http.NotFound(w, r)
	}

}