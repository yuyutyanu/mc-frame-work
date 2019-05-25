package controllers

import (
	"fmt"
	"github.com/yuyutyanu/mc-frame-work"
)

type RootController struct {
	mc_frame_work.Controller
}

func (this *RootController) Get(ct *mc_frame_work.Context){
	fmt.Println(this.ControllerName)
	fmt.Println(ct.Request.URL.Query())
	fmt.Println(ct.Params["id"])
}