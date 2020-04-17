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
	"net/http"
)

// @Router           /ping [GET]
// @Summary          Ping
// @Tag              Ping
// @ResponseEx 200   {"ping": "pong"}
func setupCommonRoute(engine *gin.Engine) {
	engine.HandleMethodNotAllowed = true
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, &gin.H{"ping": "pong"})
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
		JwtService    *service.JwtService    `di:"~"`
		CasbinService *service.CasbinService `di:"~"`
	}{}
	dic.MustInject(container)

	jwtMw := middleware.JwtMiddleware(container.JwtService)
	casbinMw := middleware.CasbinMiddleware(container.JwtService, container.CasbinService)

	v1 := engine.Group("/v1")
	{
		authCtrl := controller.NewAuthController(dic)
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/login", authCtrl.Login)
			authGroup.POST("/register", authCtrl.Register)
			authGroup.POST("/refresh", authCtrl.RefreshToken)
			authGroup.POST("/logout", jwtMw, casbinMw, authCtrl.Logout)
			authGroup.PUT("/password", jwtMw, casbinMw, authCtrl.ResetPassword)
			authGroup.GET("", jwtMw, casbinMw, authCtrl.CurrentUser)
		}

		userCtrl := controller.NewUserController(dic)
		userGroup := v1.Group("/user")
		{
			userGroup.Use(jwtMw, casbinMw)
			userGroup.GET("", userCtrl.QueryAll)
			userGroup.GET("/:uid", userCtrl.Query)
			userGroup.PUT("/:uid", userCtrl.Update)
			userGroup.DELETE("/:uid", userCtrl.Delete)
		}

		policyCtrl := controller.NewPolicyController(dic)
		policyGroup := v1.Group("/policy")
		{
			policyGroup.Use(jwtMw, casbinMw)
			policyGroup.GET("", policyCtrl.Query)
			policyGroup.POST("", policyCtrl.Insert)
			policyGroup.DELETE("", policyCtrl.Delete)
			policyGroup.PUT("/role/:uid", policyCtrl.SetRole)
		}
	}
}
