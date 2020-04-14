package server

import (
	"fmt"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/controller"
	"github.com/Aoi-hosizora/RBAC-learn/src/middleware"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/gin-gonic/gin"
	"log"
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

func setupApiRoute(engine *gin.Engine, dic *xdi.DiContainer) {
	container := &struct {
		Config *config.Config         `di:"~"`
		JwtSrv *middleware.JwtService `di:"~"`
	}{}
	if !dic.Inject(container) {
		log.Fatalf("Failed to inject")
	}

	jwtMw := container.JwtSrv.JwtMiddleware()

	userCtrl := controller.NewUserController(dic)

	v1 := engine.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/login", userCtrl.Login)
			authGroup.GET("", jwtMw, userCtrl.CurrentUser)
		}
	}
}
