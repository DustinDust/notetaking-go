package app

import (
	"net/http"
	"notetaking-go/pkg/app/config"

	"github.com/gin-gonic/gin"
)

type App struct {
	r      *gin.Engine
	router *gin.RouterGroup
	config *config.Config
}

func New(prefix string) *App {
	app := &App{}
	app.config = &config.Config{}
	app.r = gin.Default()
	app.router = app.r.Group(prefix)

	app.router.GET("ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	app.config.Init()
	return app
}

func (app *App) RunServer() {
	app.r.Run(app.config.Url)
}
