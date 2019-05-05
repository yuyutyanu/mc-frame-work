package controllers

import "fmt"

type RootController struct {
	Controller
}

func (this *RootController) Get(){
	fmt.Println(this.Ct.Params[":id"])
}