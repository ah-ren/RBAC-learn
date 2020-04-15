package middleware

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/common/exception"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

type CasbinService struct {
	Config     *config.Config `di:"~"`
	Logger     *logrus.Logger `di:"~"`
	Db         *gorm.DB       `di:"~"`
	JwtService *JwtService    `di:"~"`
}

func NewCasbinService(dic *xdi.DiContainer) *CasbinService {
	srv := &CasbinService{}
	dic.InjectForce(srv)
	return srv
}

func (b *CasbinService) CasbinMiddleware() gin.HandlerFunc {
	adapter, err := gormadapter.NewAdapterByDBUsePrefix(b.Db, "tbl_")
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		user := b.JwtService.GetContextUser(c)
		if user == nil {
			c.Abort()
			return
		}

		ok, err := b.enforce(user.Role, c.FullPath(), c.Request.Method, adapter)
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

func (b *CasbinService) enforce(sub string, obj string, act string, adapter *gormadapter.Adapter) (bool, error) {
	enforcer, err := casbin.NewEnforcer(b.Config.CasbinConfig.ConfigPath, adapter)
	if err != nil {
		return false, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return false, err
	}

	return enforcer.Enforce(sub, obj, act)
}
