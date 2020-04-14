package middleware

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/common/exception"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/qiangmzsx/string-adapter/v2"
)

type CasbinService struct {
	Config     *config.Config          `di:"~"`
	Adapter    *string_adapter.Adapter `di:"~"`
	JwtService *JwtService             `di:"~"`
}

func NewCasbinService(dic *xdi.DiContainer) *CasbinService {
	srv := &CasbinService{}
	dic.InjectForce(srv)
	return srv
}

func (b *CasbinService) CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := b.JwtService.GetContextUser(c)
		if user == nil {
			c.Abort()
			return
		}
		path := c.Request.URL.Path
		method := c.Request.Method

		ok, err := b.enforce(user.Role, path, method)
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

func (b *CasbinService) enforce(sub string, obj string, act string) (bool, error) {
	enforcer, err := casbin.NewEnforcer(b.Config.CasbinConfig.ConfigPath, b.Adapter)
	if err != nil {
		return false, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return false, err
	}

	return enforcer.Enforce(sub, obj, act)
}
