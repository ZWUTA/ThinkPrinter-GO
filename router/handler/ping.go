package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thinkprinter/models"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, models.OK.WithMsg("Pong"))
}
