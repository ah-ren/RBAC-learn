package server

import (
	"fmt"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupCommonRoute(engine *gin.Engine) {
	engine.HandleMethodNotAllowed = true
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
	engine.NoMethod(func(c *gin.Context) {
		result.Status(http.StatusMethodNotAllowed).JSON(c)
	})
	engine.NoRoute(func(c *gin.Context) {
		result.Status(http.StatusNotFound).SetMessage(fmt.Sprintf("route %s not found", c.Request.URL.Path)).JSON(c)
	})
}

func setupApiRoute(engine *gin.Engine) {

}
