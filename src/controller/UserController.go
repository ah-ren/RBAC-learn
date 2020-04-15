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
	Config      *config.ServerConfig   `di:"~"`
	Mapper      *xentity.EntityMappers `di:"~"`
	JwtService  *service.JwtService    `di:"~"`
	UserService *service.UserService   `di:"~"`
}

func NewUserController(dic *xdi.DiContainer) *UserController {
	ctrl := &UserController{}
	dic.MustInject(ctrl)
	return ctrl
}

// @Router              /v1/user [GET]
// @Summary             Query user list
// @Security            Jwt
// @Tag                 User
// @Param               page  query integer false "Query page"(default:1)
// @Param               limit query integer false "Page size" (default:10)
// @ResponseModel 200   #Result<Page<UserDto>>
func (u *UserController) QueryAll(c *gin.Context) {
	page, limit := param.BindPage(c, u.Config)
	total, users := u.UserService.QueryAll(page, limit)

	usersDto := xcondition.First(u.Mapper.MapSlice(xslice.Sti(users), &dto.UserDto{})).([]*dto.UserDto)
	result.Ok().SetPage(int32(len(usersDto)), page, total, usersDto).JSON(c)
}

// @Router              /v1/user/{uid} [GET]
// @Summary             Query user
// @Security            Jwt
// @Tag                 User
// @Param               uid path integer true "user id"
// @ResponseModel 200   #Result<UserDto>
func (u *UserController) Query(c *gin.Context) {
	id, ok := param.BindId(c, "uid")
	if !ok {
		result.Error(exception.RequestParamError).JSON(c)
		return
	}

	user, ok := u.UserService.QueryById(id)
	if !ok {
		result.Error(exception.UserNotFoundError).JSON(c)
		return
	}

	userDto := xcondition.First(u.Mapper.Map(user, &dto.UserDto{})).(*dto.UserDto)
	result.Ok().SetData(userDto).JSON(c)
}

// @Router              /v1/user/{uid} [PUT]
// @Summary             Update user
// @Security            Jwt
// @Tag                 User
// @Param               uid   path integer    true "user id"
// @Param               param body #UpdateUserParam true "request parameter"
// @ResponseModel 200   #Result
func (u *UserController) Update(c *gin.Context) {
	uid, ok := param.BindId(c, "uid")
	updateUserParam := &param.UpdateUserParam{}
	if err := c.ShouldBind(updateUserParam); err != nil || !ok {
		result.Error(exception.RequestParamError).JSON(c)
		return
	}

	user := &po.User{ID: uid}
	_ = u.Mapper.MapProp(updateUserParam, user)

	status := u.UserService.Update(user)
	if status == database.DbNotFound {
		result.Error(exception.UserNotFoundError).JSON(c)
		return
	} else if status == database.DbFailed {
		result.Error(exception.UserUpdateError).JSON(c)
		return
	}

	result.Ok().JSON(c)
}

// @Router              /v1/user/{uid} [DELETE]
// @Summary             Delete user
// @Security            Jwt
// @Tag                 User
// @Param               uid path integer true "user id"
// @ResponseModel 200   #Result
func (u *UserController) Delete(c *gin.Context) {
	id, ok := param.BindId(c, "uid")
	if !ok {
		result.Error(exception.RequestParamError).JSON(c)
		return
	}

	status := u.UserService.Delete(id)
	if status == database.DbNotFound {
		result.Error(exception.UserNotFoundError).JSON(c)
		return
	} else if status == database.DbFailed {
		result.Error(exception.UserDeleteError).JSON(c)
		return
	}

	result.Ok().JSON(c)
}
