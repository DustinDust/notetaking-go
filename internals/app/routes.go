package app

func (app *App) Routes(prefix string) {
	group := app.R.Group(prefix)
	group.GET("ping", app.Handlers.Ping)
}
