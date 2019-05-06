package main

import (
	"fmt"
	"github.com/yuyutyanu/mvc"
)

func main() {
	app := mvc.NewApp()
	str := app.Cfg.String("DB_NAME")
	fmt.Print(str)
	for _, value := range app.Cfg.GetComment(){
		fmt.Print(value)
	}

	//log2.SetLevel(log2.LevelInfo)
	//log2.Info("hogehoge")
	//log2.Debug("piyopiyo")

	//router := &mvc.ControllerRegsiter{}
	//router.Add("/user/:id/blog/:article", &controllers.RootController{})
	//err := http.ListenAndServe(":9090", router)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}