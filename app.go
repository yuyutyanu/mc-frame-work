package mvc

type App struct {
}

var StaticDir = make(map[string]string)

func (p *App) SetStaticPath(prefix string, path string) {
	StaticDir[prefix] = path
}

func NewApp() *App{
	app := new(App)
	app.SetStaticPath("/img", "./static/img")
	app.SetStaticPath("/css", "./static/css")
	return app
}

