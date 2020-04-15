package server

import (
	"fmt"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/controller"
	"github.com/Aoi-hosizora/RBAC-learn/src/middleware"
	"github.com/Aoi-hosizora/RBAC-learn/src/service"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
		Config        *config.ServerConfig   `di:"~"`
		Db            *gorm.DB               `di:"~"`
		JwtService    *service.JwtService    `di:"~"`
		CasbinService *service.CasbinService `di:"~"`
	}{}
	dic.InjectForce(container)

	jwtMw := middleware.JwtMiddleware(container.JwtService)
	casbinMw := middleware.CasbinMiddleware(container.CasbinService)

	authCtrl := controller.NewAuthController(dic)
	userCtrl := controller.NewUserController(dic)

	v1 := engine.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/login", authCtrl.Login)
			authGroup.POST("/register", authCtrl.Register)
			authGroup.POST("/refresh", authCtrl.RefreshToken)
			authGroup.POST("/logout", jwtMw, casbinMw, authCtrl.Logout)
			authGroup.PUT("/password", jwtMw, casbinMw, authCtrl.ResetPassword)
			authGroup.GET("", jwtMw, casbinMw, authCtrl.CurrentUser)
		}
		userGroup := v1.Group("/user")
		{
			userGroup.GET("", jwtMw, casbinMw, userCtrl.QueryAll)
			userGroup.GET("/:uid", jwtMw, casbinMw, userCtrl.Query)
			userGroup.PUT("/:uid", jwtMw, casbinMw, userCtrl.Update)
			userGroup.DELETE("/:uid", jwtMw, casbinMw, userCtrl.Delete)
		}
	}
}
