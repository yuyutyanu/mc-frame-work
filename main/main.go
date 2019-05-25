package main

import (
	"github.com/yuyutyanu/mc-frame-work"
	"github.com/yuyutyanu/mc-frame-work/controllers"
	"github.com/yuyutyanu/mc-frame-work/logger"
	"log"
	"net/http"
)

func main() {
	//app := mc-frame-work.NewApp()
	//str := app.Cfg.String("DB_NAME")
	//fmt.Print(str)
	//fmt.Print(app.Cfg.GetComment())

	logger.SetLevel(logger.LevelInfo)
	logger.Info("hogehoge")
	logger.Debug("piyopiyo")
	mc_frame_work.NewApp()
	router := &mc_frame_work.ControllerRegsiter{}
	router.Add("/user/:id", &controllers.RootController{})
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}