package main

import (
	"github.com/yuyutyanu/mvc"
	"github.com/yuyutyanu/mvc/controllers"
	"log"
	"net/http"
)



func main() {
	mvc.NewApp()
	router := &mvc.ControllerRegsiter{}
	router.Add("/user/:id/blog/:article", &controllers.RootController{})
	err := http.ListenAndServe(":9090", router) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}