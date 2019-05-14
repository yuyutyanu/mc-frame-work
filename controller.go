package mvc

import (
	"github.com/yuyutyanu/mvc/utils"
	"html/template"
	"net/http"
	"path"
)

type Controller struct {
	Ct *Context
	Tpl *template.Template
	Data map[interface{}]interface{}
	ControllerName string
	TplNames string
	Template []string
	TplExt string
}

type ControllerInterface interface {
	Init(childName string)
	Prepare(ct *Context)
	Get(ct *Context)
	Finish(ct *Context)
	Post(ct *Context)
	Delete(ct *Context)
	Put(ct *Context)
	Head(ct *Context)
	Patch(ct *Context)
	Options(ct *Context)
	Render(ct *Context) error
}

func (c *Controller) Init(controllerName string){
	c.Data = make(map[interface{}]interface{})
	c.Template = make([]string,0)
	c.TplNames = ""
	c.ControllerName = controllerName
	c.TplExt = "tpl"
}

func (c *Controller) Prepare(ct *Context){}
func (c *Controller) Finish(ct *Context){}

func (c *Controller) Post(ct *Context){
	http.Error(ct.ResponseWriter, "Method Not Allowed", 405)
}
func (c *Controller) Delete(ct *Context){
	http.Error(ct.ResponseWriter, "Method Not Allowed", 405)
}
func (c *Controller) Put(ct *Context){
	http.Error(ct.ResponseWriter, "Method Not Allowed", 405)
}
func (c *Controller) Head(ct *Context){
	http.Error(ct.ResponseWriter, "Method Not Allowed", 405)
}
func (c *Controller) Patch(ct *Context){
	http.Error(ct.ResponseWriter, "Method Not Allowed", 405)
}
func (c *Controller) Options(ct *Context){
	http.Error(ct.ResponseWriter, "Method Not Allowed", 405)
}

//todo 扱いづらいので base template を指定したら依存関係を解決するように
func (c *Controller) Render(ct *Context) error{
	if len(c.Template) > 0 {
		var filenames []string
		for _, file := range c.Template {
			filenames = append(filenames, path.Join("./views/template/", file))
		}
		t, err := template.ParseFiles(filenames...)
		utils.DoError(err)
		err = t.ExecuteTemplate(ct.ResponseWriter, c.TplNames, c.Data)
		utils.DoError(err)
	} else {
		if c.TplNames == "" {
			c.TplNames = c.ControllerName + "/" + ct.Request.Method + "." + c.TplExt
		}
		t, err := template.ParseFiles(path.Join("./views/", c.TplNames))
		utils.DoError(err)
		err = t.Execute(ct.ResponseWriter, c.Data)
		utils.DoError(err)
	}
	return nil
}
//func (c *Controller) Redirect(url string, code int) {
//	c.Ct.Redirect(code, url)
//}


