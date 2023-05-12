package app

import (
	"notetaking-go/internals/app/config"
	"notetaking-go/internals/app/db"
	"notetaking-go/internals/app/handlers"
	"notetaking-go/internals/app/loggers"

	"github.com/gin-gonic/gin"
)

type App struct {
	R        *gin.Engine
	Handlers *handlers.Handlers
	Config   *config.Config
	DB       *db.Database
	Logger   *loggers.Logger
}

func New(prefix string, configFile string) *App {
	app := &App{}
	app.R = gin.Default()
	app.Config = &config.Config{}
	app.DB = &db.Database{}
	app.Handlers = &handlers.Handlers{}
	app.Logger = &loggers.Logger{}

	app.Config.Init(configFile)
	app.Logger.Init(loggers.DEVELOPMENT)
	app.DB.StartDBConnection(app.Config.DB)
	app.DB.Seeds()
	app.Routes(prefix)

	return app
}

func (app *App) RunServer() {
	app.R.Run(app.Config.Url)
}
