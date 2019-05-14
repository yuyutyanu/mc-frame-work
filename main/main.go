package main

import (
	"github.com/yuyutyanu/mvc"
	"github.com/yuyutyanu/mvc/controllers"
	"log"
	"net/http"
)

func main() {
	//app := mvc.NewApp()
	//str := app.Cfg.String("DB_NAME")
	//fmt.Print(str)
	//for _, value := range app.Cfg.GetComment(){
	//	fmt.Print(value)
	//}

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

	//s1, _ := template.ParseFiles("./views/template/header.gtpl", "./views/template/content.gtpl", "./views/template/footer.gtpl")
	//fmt.Println()
	//s1.ExecuteTemplate(os.Stdout, "content", nil)
	//fmt.Println()
	//s1.Execute(os.Stdout, nil)
}