package mvc

type App struct {
	Cfg *Config
}

var StaticDir = make(map[string]string)

func (p *App) SetStaticPath(prefix string, path string) {
	StaticDir[prefix] = path
}

func NewApp() *App{
	app := new(App)
	app.SetStaticPath("/img", "./static/img")
	app.SetStaticPath("/css", "./static/css")
	app.SetStaticPath("/js", "./static/js")
	cfg, err := LoadConfig(".env")
	if err != nil{
		panic(err)
	}
	app.Cfg = cfg
	return app
}

