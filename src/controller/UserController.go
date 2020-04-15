package controller

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/common/exception"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/database"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/dto"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/param"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/po"
	"github.com/Aoi-hosizora/RBAC-learn/src/service"
	"github.com/Aoi-hosizora/ahlib/xcondition"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xentity"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Config     *config.ServerConfig    `di:"~"`
	Mapper     *xentity.EntityMappers  `di:"~"`
	JwtService *service.JwtService     `di:"~"`
	UserRepo   *service.UserRepository `di:"~"`
}

func NewUserController(dic *xdi.DiContainer) *UserController {
	ctrl := &UserController{}
	dic.InjectForce(ctrl)
	return ctrl
}

func (u *UserController) QueryAll(c *gin.Context) {
	page, limit := param.BindPage(c, u.Config)
	total, users := u.UserRepo.QueryAll(page, limit)

	usersDto := xcondition.First(u.Mapper.MapSlice(xslice.Sti(users), &dto.UserDto{})).([]*dto.UserDto)
	result.Ok().SetPage(int32(len(usersDto)), page, total, usersDto).JSON(c)
}

func (u *UserController) Query(c *gin.Context) {
	id, ok := param.BindId(c, "uid")
	if !ok {
		result.Error(exception.RequestParamError).JSON(c)
		return
	}

	user := u.UserRepo.QueryById(id)
	if user == nil {
		result.Error(exception.UserNotFoundError).JSON(c)
		return
	}

	userDto := xcondition.First(u.Mapper.Map(user, &dto.UserDto{})).(*dto.UserDto)
	result.Ok().SetData(userDto).JSON(c)
}

func (u *UserController) Update(c *gin.Context) {
	uid, ok := param.BindId(c, "uid")
	userParam := &param.UserParam{}
	if err := c.ShouldBind(userParam); err != nil || !ok {
		result.Error(exception.RequestParamError).JSON(c)
		return
	}

	user := &po.User{ID: uid}
	_ = u.Mapper.MapProp(userParam, user)

	status := u.UserRepo.Update(user)
	if status == database.DbNotFound {
		result.Error(exception.UserNotFoundError).JSON(c)
		return
	} else if status == database.DbFailed {
		result.Error(exception.UserUpdateError).JSON(c)
		return
	}

	userDto := xcondition.First(u.Mapper.Map(user, &dto.UserDto{})).(*dto.UserDto)
	result.Ok().SetData(userDto).JSON(c)
}

func (u *UserController) Delete(c *gin.Context) {
	id, ok := param.BindId(c, "uid")
	if !ok {
		result.Error(exception.RequestParamError).JSON(c)
		return
	}

	status := u.UserRepo.Delete(id)
	if status == database.DbNotFound {
		result.Error(exception.UserNotFoundError).JSON(c)
		return
	} else if status == database.DbFailed {
		result.Error(exception.UserDeleteError).JSON(c)
		return
	}

	result.Ok().JSON(c)
}
