package mc_frame_work

import (
	"html/template"
	"net/http"
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

//func (c *Controller) Redirect(url string, code int) {
//	c.Ct.Redirect(code, url)
//}


