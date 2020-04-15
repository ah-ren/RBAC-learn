package service

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type CasbinService struct {
	Config     *config.ServerConfig `di:"~"`
	Logger     *logrus.Logger       `di:"~"`
	Db         *gorm.DB             `di:"~"`
	JwtService *JwtService          `di:"~"`
}

func NewCasbinService(dic *xdi.DiContainer) *CasbinService {
	srv := &CasbinService{}
	dic.InjectForce(srv)
	return srv
}

func (c *CasbinService) GetAdapter() *gormadapter.Adapter {
	adapter, err := gormadapter.NewAdapterByDBUsePrefix(c.Db, "tbl_")
	if err != nil {
		panic(err)
	}
	return adapter
}

func (c *CasbinService) Enforce(sub string, obj string, act string, adapter *gormadapter.Adapter) (bool, error) {
	enforcer, err := casbin.NewEnforcer(c.Config.CasbinConfig.ConfigPath, adapter)
	if err != nil {
		return false, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return false, err
	}

	return enforcer.Enforce(sub, obj, act)
}

func (c *CasbinService) GetContextRole(context *gin.Context) (string, bool) {
	user := c.JwtService.GetContextUser(context)
	if user == nil {
		return "", false
	}
	return user.Role, true
}
