package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"thinkprinter/initializer"
	"thinkprinter/models"
	"thinkprinter/router"
)

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
	initializer.PreLaunch()
}
