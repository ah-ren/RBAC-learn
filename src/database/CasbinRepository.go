package database

import (
	"fmt"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/po"
)

type CasbinRepository struct {
	_list []*po.Casbin
}

func NewCasbinRepository() *CasbinRepository {
	return &CasbinRepository{_list: []*po.Casbin{
		{ID: 1, PType: "p", Role: "admin", Path: "/v1/user", Method: "GET"},
		{ID: 2, PType: "p", Role: "admin", Path: "/v1/auth", Method: "GET"},
		{ID: 3, PType: "p", Role: "normal", Path: "/v1/auth", Method: "GET"},
	}}
}

func (c *CasbinRepository) String() string {
	out := ""
	for _, o := range c._list {
		out += fmt.Sprintf("%s, %s, %s, %s\n", o.PType, o.Role, o.Path, o.Method)
	}
	return out
}
