package app

import (
	"notetaking-go/internals/app/config"
	"notetaking-go/internals/app/handlers"
	"notetaking-go/internals/app/loggers"
	"notetaking-go/internals/app/models"

	"github.com/gin-gonic/gin"
)

type App struct {
	R        *gin.Engine
	Handlers *handlers.Handlers
	Config   *config.Config
	Models   *models.Models
	Logger   *loggers.Logger
}

func New(prefix string, configFile string, dbUrl string) *App {
	app := &App{}
	app.R = gin.Default()
	app.Config = &config.Config{}
	app.Models = &models.Models{}
	app.Handlers = &handlers.Handlers{}
	app.Logger = &loggers.Logger{}

	{
		app.Config.Init(configFile)
		app.Logger.Init(loggers.DEVELOPMENT)
	}
	{
		app.Models.StartDBConnection(dbUrl)
		app.Models.Seeds()
	}
	{
		app.Routes(prefix)
	}

	return app
}

func (app *App) RunServer() {
	app.R.Run(app.Config.Url)
}
