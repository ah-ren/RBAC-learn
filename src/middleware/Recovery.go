package middleware

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/common/exception"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/ahlib/xcolor"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	// 0: exception.stack
	// 1: exception.NewErrorDto
	// 2: middleware.RecoveryMiddleware.func1.1
	// 3: runtime.gopanic
	// 4: server.setupCommonRoute.func3
	// 5: gin.(*Context).Next
	return func(c *gin.Context) {
		// skip, _ := strconv.Atoi(c.Query("s"))
		skip := 4 // stack
		gin.Recovery()
		defer func() {
			if err := recover(); err != nil {
				r := result.Error(exception.ServerRecoveryError)

				if gin.Mode() == gin.DebugMode {
					errDto := exception.NewErrorDto(err, skip, c, true)
					r.Error = errDto
				}
				r.JSON(c)

				logger.Debugln("!!!!!!!!!!!!!!!!!!")
				logger.Errorf(xcolor.Red.Paint("[Recovery] panic recovered: %s"), err)
				logger.Debugln("!!!!!!!!!!!!!!!!!!")
			}
		}()
		c.Next()
	}
}
