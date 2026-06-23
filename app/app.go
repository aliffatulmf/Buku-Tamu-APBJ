package app

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type App struct {
	Server *gin.Engine
	Port   string
}

func New(server *gin.Engine) *App {
	return &App{Server: server}
}

func (app *App) SetPort(port string) {
	if strings.Contains(port, ":") {
		app.Port = port
	} else {
		app.Port = fmt.Sprintf(":%s", port)
	}
}

func (app *App) SetHandler(f func(a *App)) {
	f(app)
}

func (app *App) RunHTTP() {
	if err := app.Server.Run(app.Port); err != nil {
		fmt.Println("FATAL:", err.Error())
	}
}
