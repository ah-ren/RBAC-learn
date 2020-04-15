package middleware

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/RBAC-learn/src/service"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware(srv *service.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := srv.GetToken(c)
		user, err := srv.JwtCheck(token)
		if err != nil {
			result.Error(err).JSON(c)
			c.Abort()
			return
		}

		c.Set(srv.UserKey, user)
		c.Next()
	}
}
