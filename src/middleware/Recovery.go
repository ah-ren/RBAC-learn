package middleware

import (
	"fmt"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/exception"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/ahlib/xcolor"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http/httputil"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func RecoveryMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// skip, _ := strconv.Atoi(c.Query("s"))
		skip := 2 // stack
		defer func() {
			if err := recover(); err != nil {
				r := result.Error(exception.ServerRecoveryError)

				if gin.Mode() == gin.DebugMode {
					now := time.Now().Format(time.RFC3339)

					// runtime
					function := "?"
					pc, filename, line, ok := runtime.Caller(skip)
					if ok {
						function = runtime.FuncForPC(pc).Name()
						_, function = filepath.Split(function)
					}
					detail := fmt.Sprintf("%v", err)

					// request
					requestBytes, _ := httputil.DumpRequest(c.Request, false)
					requestParams := strings.Split(string(requestBytes), "\r\n")
					request := make([]string, 0)
					for _, param := range requestParams {
						if strings.HasPrefix(param, "Authorization:") { // Authorization header
							request = append(request, "Authorization: *")
						} else if param != "" { // other param
							request = append(request, param)
						}
					}

					// response
					r.SetData(&exception.ErrorDto{
						Time:    now,
						File:    filename,
						Line:    line,
						Func:    function,
						Detail:  detail,
						Request: request,
					})
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
