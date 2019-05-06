package mvc

import "github.com/yuyutyanu/mvc/conf"

type App struct {
	Cfg *conf.Config
}

var StaticDir = make(map[string]string)

func (p *App) SetStaticPath(prefix string, path string) {
	StaticDir[prefix] = path
}

func NewApp() *App{
	app := new(App)
	app.SetStaticPath("/img", "./static/img")
	app.SetStaticPath("/css", "./static/css")
	cfg, err := conf.LoadConfig(".env")
	if err != nil{
		panic(err)
	}
	app.Cfg = cfg
	return app
}

