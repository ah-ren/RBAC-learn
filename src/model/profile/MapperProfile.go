package profile

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/model/dto"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/po"
	"github.com/Aoi-hosizora/ahlib/xentity"
)

func CreateEntityMappers() *xentity.EntityMappers {
	mappers := xentity.NewEntityMappers()

	mappers.AddMapper(xentity.NewEntityMapper(&po.User{}, func() interface{} { return &dto.UserDto{} }, func(from interface{}, to interface{}) error {
		user := from.(*po.User)
		userDto := to.(*dto.UserDto)

		userDto.ID = user.ID
		userDto.Name = user.Name
		userDto.Role = user.Role
		return nil
	}))

	return mappers
}
