package controllers

import (
	"fmt"
	"github.com/yuyutyanu/mvc"
)

type RootController struct {
	mvc.Controller
}

func (this *RootController) Get(ct *mvc.Context){
	fmt.Println(this.ControllerName)
	fmt.Println(ct.Request.URL.Query())
	fmt.Println(ct.Params["id"])
}