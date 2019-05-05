package main

import (
	"github.com/yuyutyanu/mvc/controllers"
	"log"
	"net/http"
)



func main() {
	handler := &controllers.ControllerRegsiter{}
	handler.Add("/user/:id/blog/:article", &controllers.RootController{})
	err := http.ListenAndServe(":9090", handler) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}