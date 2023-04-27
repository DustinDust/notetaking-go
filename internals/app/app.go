package app

import (
	"notetaking-go/internals/app/config"
	"notetaking-go/internals/app/handlers"
	"notetaking-go/internals/app/models"

	"github.com/gin-gonic/gin"
)

type App struct {
	R        *gin.Engine
	Router   *gin.RouterGroup
	Handlers *handlers.Handler
	Config   *config.Config
	Models   *models.Models
}

func New(prefix string, configFile string, dbUrl string) *App {
	app := &App{}
	app.R = gin.Default()
	app.Config = &config.Config{}
	app.Models = &models.Models{}
	app.Handlers = &handlers.Handler{}
	app.Router = app.R.Group(prefix)

	{
		app.Config.Init(configFile)
	}

	{
		app.Models.StartDBConnection(dbUrl)
		app.Models.Seeds()
	}
	return app
}

func (app *App) RunServer() {
	app.R.Run(app.Config.Url)
}
