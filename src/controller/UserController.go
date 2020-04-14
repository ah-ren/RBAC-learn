package controller

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/database"
	"github.com/Aoi-hosizora/RBAC-learn/src/middleware"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/dto"
	"github.com/Aoi-hosizora/ahlib/xcondition"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xentity"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Config     *config.Config           `di:"~"`
	JwtService *middleware.JwtService   `di:"~"`
	Mapper     *xentity.EntityMappers   `di:"~"`
	UserRepo   *database.UserRepository `di:"~"`
}

func NewUserController(dic *xdi.DiContainer) *UserController {
	ctrl := &UserController{}
	dic.InjectForce(ctrl)
	return ctrl
}

func (u *UserController) QueryAll(c *gin.Context) {
	users := u.UserRepo.QueryAll()
	usersDto := xcondition.First(u.Mapper.MapSlice(xslice.Sti(users), &dto.UserDto{})).([]*dto.UserDto)
	result.Ok().SetPage(int32(len(usersDto)), 1, 200, usersDto)
}
