package database

import (
	"fmt"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/po"
	"github.com/qiangmzsx/string-adapter/v2"
)

type CasbinRepository struct {
	_list []*po.Casbin
}

func NewCasbinRepository() *CasbinRepository {
	return &CasbinRepository{_list: []*po.Casbin{
		{ID: 1, PType: "p", Role: "Admin", Path: "/v1/user", Method: "GET"},
		{ID: 2, PType: "p", Role: "Admin", Path: "/v1/auth", Method: "GET"},
		{ID: 3, PType: "p", Role: "Normal", Path: "/v1/auth", Method: "GET"},
	}}
}

func (c *CasbinRepository) String() string {
	out := ""
	for _, o := range c._list {
		out += fmt.Sprintf("%s, %s, %s, %s\n", o.PType, o.Role, o.Path, o.Method)
	}
	return out
}

func (c *CasbinRepository) Adapter() *string_adapter.Adapter {
	return string_adapter.NewAdapter(c.String())
}
