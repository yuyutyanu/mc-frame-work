package controllers

import (
	"fmt"
	"github.com/yuyutyanu/mvc"
)

type RootController struct {
	mvc.Controller
}

func (this *RootController) Get(){
	fmt.Println(this.Ct.Params[":id"])
}