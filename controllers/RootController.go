package controllers

import (
	"github.com/yuyutyanu/mvc"
)

type RootController struct {
	mvc.Controller
}

func (this *RootController) Get(){
	this.Template = append(this.Template, "footer.gtpl")
	this.Template = append(this.Template, "header.gtpl")
	this.Template = append(this.Template, "content.gtpl")
	this.TplNames = "content"
}