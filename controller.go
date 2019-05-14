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
	ChildName string
	TplNames string
	Template []string
	TplExt string
}

type ControllerInterface interface {
	Init(ct *Context, cn string)
	Prepare()
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	Finish()
	Render() error
}

func (c *Controller) Init(ct *Context, controllerName string){
	c.Data = make(map[interface{}]interface{})
	c.Template = make([]string,0)
	c.TplNames = ""
	c.ChildName = controllerName
	c.Ct = ct
	c.TplExt = "tpl"
}

func (c *Controller) Prepare(){}
func (c *Controller) Finish(){}

func (c *Controller) Post(){
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}
func (c *Controller) Delete(){
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}
func (c *Controller) Put(){
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}
func (c *Controller) Head(){
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}
func (c *Controller) Patch(){
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}
func (c *Controller) Options(){
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Render() error{
	if len(c.Template) > 0 {
		var filenames []string
		for _, file := range c.Template {
			filenames = append(filenames, path.Join("./views/template/", file))
		}
		t, err := template.ParseFiles(filenames...)
		utils.DoError(err)
		err = t.ExecuteTemplate(c.Ct.ResponseWriter, c.TplNames, c.Data)
		utils.DoError(err)
	} else {
		if c.TplNames == "" {
			c.TplNames = c.ChildName + "/" + c.Ct.Request.Method + "." + c.TplExt
		}
		t, err := template.ParseFiles(path.Join("./views/", c.TplNames))
		utils.DoError(err)
		err = t.Execute(c.Ct.ResponseWriter, c.Data)
		utils.DoError(err)
	}
	return nil
}
//func (c *Controller) Redirect(url string, code int) {
//	c.Ct.Redirect(code, url)
//}


