package web

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

var F embed.FS

func Index(c *gin.Context) {
	c.FileFromFS("static/", http.FS(F))
}
