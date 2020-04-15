package middleware

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/common/exception"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/RBAC-learn/src/service"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CasbinMiddleware(srv *service.CasbinService) gin.HandlerFunc {
	adapter := srv.GetAdapter()
	return func(c *gin.Context) {
		sub, ok := srv.GetContextRole(c)
		if !ok {
			c.Abort()
			result.Error(exception.CheckUserRoleError).JSON(c)
			return
		}
		obj := c.FullPath()
		act := c.Request.Method

		ok, err := srv.Enforce(sub, obj, act, adapter)
		if err != nil {
			c.Abort()
			result.Error(exception.CheckUserRoleError).JSON(c)
			return
		}
		if !ok {
			c.Abort()
			result.Error(exception.RoleHasNoPermissionError).JSON(c)
			return
		}
		c.Next()
	}
}
