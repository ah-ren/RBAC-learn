package controller

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/common/exception"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/database"
	"github.com/Aoi-hosizora/RBAC-learn/src/middleware"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/dto"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/param"
	"github.com/Aoi-hosizora/RBAC-learn/src/util"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
)

type AuthController struct {
	Config     *config.Config           `di:"~"`
	Logger     *logrus.Logger           `di:"~"`
	JwtService *middleware.JwtService   `di:"~"`
	UserRepo   *database.UserRepository `di:"~"`
}

func NewUserController(dic *xdi.DiContainer) *AuthController {
	ctrl := &AuthController{}
	if !dic.Inject(ctrl) {
		log.Fatalf("Failed to inject")
	}
	return ctrl
}

func (a *AuthController) Login(c *gin.Context) {
	loginParam := &param.LoginParam{}
	if err := c.ShouldBind(loginParam); err != nil {
		result.Error(exception.RequestParamError).JSON(c)
		return
	}
	user := a.UserRepo.QueryById(loginParam.Id)
	if user.Password != loginParam.Password {
		result.Error(exception.WrongPasswordError).JSON(c)
		return
	}

	token, err := util.AuthUtil.CreateToken(user.ID, a.Config.JwtConfig)
	if err != nil {
		result.Error(exception.LoginError).SetData(err).JSON(c)
	}

	loginDto := dto.LoginDto{
		User: &dto.UserDto{
			ID:   user.ID,
			Name: user.Name,
			Role: user.Role,
		},
		Token: token,
	}
	result.Ok().SetData(loginDto).JSON(c)
}

func (a *AuthController) CurrentUser(c *gin.Context) {
	user := a.JwtService.GetContextUser(c)
	userDto := &dto.UserDto{
		ID:   user.ID,
		Name: user.Name,
		Role: user.Role,
	}
	result.Ok().SetData(userDto).JSON(c)
}
