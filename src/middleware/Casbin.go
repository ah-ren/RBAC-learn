package middleware

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/common/exception"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/RBAC-learn/src/service"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CasbinMiddleware(jwtService *service.JwtService, casbinService *service.CasbinService) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := jwtService.GetContextUser(c)
		if user == nil {
			c.Abort()
			result.Error(exception.CheckUserRoleError).JSON(c)
			return
		}
		sub := user.Role
		obj := c.FullPath()
		act := c.Request.Method

		ok, err := casbinService.Enforce(sub, obj, act)
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
