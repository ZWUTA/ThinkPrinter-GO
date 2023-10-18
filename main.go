package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"thinkprinter/initializer"
	"thinkprinter/models"
	"thinkprinter/router"
	"thinkprinter/router/handler"
)

//go:embed static
var f embed.FS

func main() {
	url := fmt.Sprintf("%s:%d", models.C.Core.Bind, models.C.Core.Port)
	slog.Info("ThinkPrinter 正在监听地址", "url", url)

	gin.SetMode(gin.ReleaseMode)

	r := router.SetupRouter()
	err := r.Run(url)
	if err != nil {
		panic(err)
	}
}

func init() {
	handler.F = f
	initializer.PreLaunch()
}
