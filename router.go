package mvc

import (
	"github.com/yuyutyanu/mvc/utils"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type controllerInfo struct {
	regex      *regexp.Regexp
	params     map[int]string
	controller ControllerInterface
}
type ControllerRegsiter struct {
	routes []*controllerInfo
}

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         map[string]string
}

func (p *ControllerRegsiter) Add(pettern string, c ControllerInterface) {
	parts := strings.Split(pettern, "/")
	params := make(map[int]string)

	j := 0
	for i, part := range parts {
		// '/part/:param'
		if strings.HasPrefix(part, ":") {
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
	utils.DoError(regexErr) //TODO add error handling here to avoid panic

	route := &controllerInfo{}
	route.regex = regex
	route.params = params
	route.controller = c

	p.routes = append(p.routes, route)
}

func (p *ControllerRegsiter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			if err != nil {
				panic(err) //TODO add error handling here to avoid panic
			}
			//if !RecoverPanic {
			//	panic(err)
			//}else {
			//	log.Critical("Handler crashed with error", err)
			//	for i := 1; ;i +=1  {
			//		_, file, line,ok := runtime.Caller(i)
			//		if !ok {
			//			break;
			//		}
			//		log.Critical(file, line)
			//	}
			//}
		}
	}()

	// static routing
	var started bool
	requestPath := r.URL.Path
	for prefix, staticDir := range StaticDir {
		if strings.HasPrefix(requestPath, prefix) {
			file := staticDir + requestPath[len(prefix):]
			http.ServeFile(w, r, file)
			started = true
			return
		}
	}

	for _, route := range p.routes {
		if !route.regex.MatchString(requestPath) {
			continue
		}

		// /:id..etc params join query string
		matches := route.regex.FindStringSubmatch(requestPath)
		if len(matches[0]) != len(requestPath) {
			continue
		}

		params := make(map[string]string)
		if len(route.params) > 0 {
			for i, match := range matches[1:] {
				key := strings.Replace(route.params[i], ":", "", -1)
				params[key] = match
			}
		}

		ct := &Context{ResponseWriter: w, Request: r, Params: params}
		childName := reflect.Indirect(reflect.ValueOf(route.controller)).Type().Name()
		route.controller.Init(childName)
		route.controller.Prepare(ct)

		// if　はださいので　route.getみたいに叩ける用変える
		if r.Method == "GET" {
			route.controller.Get(ct)
		} else if r.Method == "POST" {
			route.controller.Post(ct)
		} else if r.Method == "HEAD" {
			route.controller.Head(ct)
		} else if r.Method == "DELETE" {
			route.controller.Delete(ct)
		} else if r.Method == "PUT" {
			route.controller.Put(ct)
		} else if r.Method == "PATCH" {
			route.controller.Patch(ct)
		} else if r.Method == "OPTIONS" {
			route.controller.Options(ct)
		}

		route.controller.Finish(ct)
		started = true
		break
	}

	if started == false {
		http.NotFound(w, r)
	}

}
