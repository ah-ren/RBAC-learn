package service

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/database"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/dto"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type CasbinService struct {
	Config     *config.ServerConfig `di:"~"`
	Logger     *logrus.Logger       `di:"~"`
	Db         *gorm.DB             `di:"~"`
	JwtService *JwtService          `di:"~"`

	Adapter *gormadapter.Adapter `di:"-"`
}

func NewCasbinService(dic *xdi.DiContainer) *CasbinService {
	srv := &CasbinService{}
	dic.MustInject(srv)

	adapter, err := gormadapter.NewAdapterByDBUsePrefix(srv.Db, "tbl_")
	if err != nil {
		panic(err)
	}
	srv.Adapter = adapter
	return srv
}

func (c *CasbinService) GetEnforcer() (*casbin.Enforcer, error) {
	enforcer, err := casbin.NewEnforcer(c.Config.CasbinConfig.ConfigPath, c.Adapter)
	if err != nil {
		return nil, err
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return enforcer, nil
}

func (c *CasbinService) Enforce(sub string, obj string, act string) (bool, error) {
	enforcer, err := c.GetEnforcer()
	if err != nil {
		return false, nil
	}
	return enforcer.Enforce(sub, obj, act)
}

func (c *CasbinService) GetRoles() ([]string, bool) {
	enforcer, err := c.GetEnforcer()
	if err != nil {
		return nil, false
	}
	return enforcer.GetAllRoles(), true
}

func (c *CasbinService) GetPolicies() ([]*dto.PolicyDto, bool) {
	enforcer, err := c.GetEnforcer()
	if err != nil {
		return nil, false
	}
	policies := enforcer.GetPolicy()
	out := make([]*dto.PolicyDto, len(policies))
	for idx := range policies {
		if len(policies[idx]) < 3 {
			continue
		}
		out[idx] = &dto.PolicyDto{
			Role:   policies[idx][0],
			Path:   policies[idx][1],
			Method: policies[idx][2],
		}
	}
	return out, true
}

func (c *CasbinService) AddPolicy(sub string, obj string, act string) database.DbStatus {
	enforcer, err := c.GetEnforcer()
	if err != nil {
		return database.DbFailed
	}
	ok, err := enforcer.AddPolicy(sub, obj, act)
	if !ok {
		return database.DbExisted
	} else if err != nil {
		return database.DbFailed
	}
	return database.DbSuccess
}

func (c *CasbinService) DeletePolicy(sub string, obj string, act string) database.DbStatus {
	enforcer, err := c.GetEnforcer()
	if err != nil {
		return database.DbFailed
	}
	ok, err := enforcer.RemovePolicy(sub, obj, act)
	if !ok {
		return database.DbNotFound
	} else if err != nil {
		return database.DbFailed
	}
	return database.DbSuccess
}
