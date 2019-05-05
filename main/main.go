package main

import (
	"github.com/yuyutyanu/mvc"
	"github.com/yuyutyanu/mvc/controllers"
	log2 "github.com/yuyutyanu/mvc/log"
	"log"
	"net/http"
)



func main() {
	log2.SetLevel(log2.LevelInfo)
	log2.Info("hogehoge")
	log2.Debug("piyopiyo")

	mvc.NewApp()
	router := &mvc.ControllerRegsiter{}
	router.Add("/user/:id/blog/:article", &controllers.RootController{})
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}