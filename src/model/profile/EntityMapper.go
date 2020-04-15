package profile

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/model/dto"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/param"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/po"
	"github.com/Aoi-hosizora/ahlib/xentity"
)

func CreateEntityMappers() *xentity.EntityMappers {
	mappers := xentity.NewEntityMappers()

	// userPo -> userDto
	mappers.AddMapper(xentity.NewEntityMapper(&po.User{}, func() interface{} { return &dto.UserDto{} }, func(from interface{}, to interface{}) error {
		user := from.(*po.User)
		userDto := to.(*dto.UserDto)

		userDto.ID = user.ID
		userDto.Name = user.Name
		userDto.Role = user.Role
		return nil
	}))

	// registerParam -> userPo
	mappers.AddMapper(xentity.NewEntityMapper(&param.RegisterParam{}, func() interface{} { return &po.User{} }, func(from interface{}, to interface{}) error {
		registerParam := from.(*param.RegisterParam)
		user := to.(*po.User)

		user.Password = registerParam.Password
		user.Name = registerParam.Name
		return nil
	}))

	// updateUserParam -> userPo
	mappers.AddMapper(xentity.NewEntityMapper(&param.UpdateUserParam{}, func() interface{} { return &po.User{} }, func(from interface{}, to interface{}) error {
		userParam := from.(*param.UpdateUserParam)
		user := to.(*po.User)

		user.Name = userParam.Name
		return nil
	}))

	return mappers
}
