package main

import (
	"fmt"
	"github.com/yuyutyanu/mvc"
	"github.com/yuyutyanu/mvc/controllers"
	"log"
	"net/http"
)

func main() {
	app := mvc.NewApp()
	//str := app.Cfg.String("DB_NAME")
	//fmt.Print(str)
		fmt.Print(app.Cfg.GetComment())

	//log2.SetLevel(log2.LevelInfo)
	//log2.Info("hogehoge")
	//log2.Debug("piyopiyo")
	mvc.NewApp()
	router := &mvc.ControllerRegsiter{}
	router.Add("/user/:id", &controllers.RootController{})
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}