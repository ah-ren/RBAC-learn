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
	"github.com/Aoi-hosizora/RBAC-learn/src/util"
	"github.com/Aoi-hosizora/ahlib/xcondition"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xentity"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Config       *config.ServerConfig   `di:"~"`
	Mapper       *xentity.EntityMappers `di:"~"`
	JwtService   *service.JwtService    `di:"~"`
	UserService  *service.UserService   `di:"~"`
	TokenService *service.TokenService  `di:"~"`
}

func NewAuthController(dic *xdi.DiContainer) *AuthController {
	ctrl := &AuthController{}
	dic.MustInject(ctrl)
	return ctrl
}

// @Router              /v1/auth/login [POST]
// @Summary             Login
// @Tag                 Authorization
// @Param               param body #LoginParam true "Request parameter"
// @ResponseModel 200   #Result<LoginDto>
func (a *AuthController) Login(c *gin.Context) {
	loginParam := &param.LoginParam{}
	if err := c.ShouldBind(loginParam); err != nil {
		result.Error(exception.RequestParamError).JSON(c)
		return
	}

	user, ok := a.UserService.QueryById(loginParam.Id)
	if !ok {
		result.Error(exception.UserNotFoundError).JSON(c)
		return
	}

	if ok, err := util.AuthUtil.CheckPassword(loginParam.Password, user.Password); err != nil {
		result.Error(exception.LoginError).JSON(c)
		return
	} else if !ok {
		result.Error(exception.WrongPasswordError).JSON(c)
		return
	}

	accessToken, err := util.AuthUtil.GenerateToken(user.ID, false, a.Config.JwtConfig)
	if err != nil {
		result.Error(exception.LoginError).JSON(c)
		return
	}
	refreshToken, err := util.AuthUtil.GenerateToken(user.ID, true, a.Config.JwtConfig)
	if err != nil {
		result.Error(exception.LoginError).JSON(c)
		return
	}
	ok = a.TokenService.Insert(accessToken, user.ID)
	if !ok {
		result.Error(exception.LoginError).JSON(c)
		return
	}

	userDto := xcondition.First(a.Mapper.Map(user, &dto.UserDto{})).(*dto.UserDto)
	result.Ok().SetData(&dto.LoginDto{
		User:         userDto,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}).JSON(c)
}

// @Router              /v1/auth/register [POST]
// @Summary             Register
// @Tag                 Authorization
// @Param               param body #RegisterParam true "Request parameter"
// @ResponseModel 200   #Result<UserDto>
func (a *AuthController) Register(c *gin.Context) {
	registerParam := &param.RegisterParam{}
	if err := c.ShouldBind(registerParam); err != nil {
		result.Error(exception.RequestParamError).JSON(c)
		return
	}

	user := &po.User{}
	_ = a.Mapper.MapProp(registerParam, user)
	newPassword, err := util.AuthUtil.EncryptPassword(user.Password)
	if err != nil {
		result.Error(exception.RegisterError).JSON(c)
		return
	}

	user.Password = newPassword
	status := a.UserService.Insert(user)
	if status == database.DbFailed {
		result.Error(exception.RegisterError).JSON(c)
		return
	}

	userDto := xcondition.First(a.Mapper.Map(user, &dto.UserDto{})).(*dto.UserDto)
	result.Ok().SetData(userDto).JSON(c)
}

// @Router              /v1/auth/token [POST]
// @Summary             Refresh token
// @Tag                 Authorization
// @Param               param body #TokenParam true "Request parameter"
// @ResponseModel 200   #Result<TokenDto>
func (a *AuthController) RefreshToken(c *gin.Context) {
	tokenParam := &param.TokenParam{}
	if err := c.ShouldBind(tokenParam); err != nil {
		result.Error(exception.RequestParamError).JSON(c)
		return
	}

	uid, refreshToken, accessToken, err := util.AuthUtil.RefreshToken(tokenParam.RefreshToken, tokenParam.AccessToken, a.Config.JwtConfig)
	if err != nil {
		if err == exception.InvalidRefreshTokenError {
			result.Error(exception.InvalidRefreshTokenError).JSON(c)
		} else {
			result.Error(exception.RefreshTokenError).JSON(c)
		}
		return
	}
	ok := a.TokenService.Insert(accessToken, uid)
	if !ok {
		result.Error(exception.RefreshTokenError).JSON(c)
		return
	}

	result.Ok().SetData(&dto.TokenDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}).JSON(c)
}

// @Router              /v1/auth [GET]
// @Summary             Current user
// @Security            Jwt
// @Tag                 Authorization
// @ResponseModel 200   #Result<UserDto>
func (a *AuthController) CurrentUser(c *gin.Context) {
	user := a.JwtService.GetContextUser(c)

	userDto := xcondition.First(a.Mapper.Map(user, &dto.UserDto{})).(*dto.UserDto)
	result.Ok().SetData(userDto).JSON(c)
}

// @Router              /v1/auth/logout [POST]
// @Summary             Logout
// @Security            Jwt
// @Tag                 Authorization
// @ResponseModel 200   #Result
func (a *AuthController) Logout(c *gin.Context) {
	token := a.JwtService.GetToken(c)
	ok := a.TokenService.Delete(token)
	if !ok {
		result.Error(exception.LogoutError).JSON(c)
		return
	}

	result.Ok().JSON(c)
}

// @Router              /v1/auth/password [PUT]
// @Summary             Reset password
// @Security            Jwt
// @Tag                 Authorization
// @Param               param body #PasswordParam true "Request parameter"
// @ResponseModel 200   #Result
func (a *AuthController) ResetPassword(c *gin.Context) {
	user := a.JwtService.GetContextUser(c)
	passParam := &param.PasswordParam{}
	if err := c.ShouldBind(passParam); err != nil {
		result.Error(exception.RequestParamError).JSON(c)
		return
	}

	encrypted, err := util.AuthUtil.EncryptPassword(passParam.Password)
	if err != nil {
		result.Error(exception.UpdatePasswordError).JSON(c)
		return
	}
	user.Password = encrypted

	status := a.UserService.Update(user)
	if status == database.DbNotFound {
		result.Error(exception.UserNotFoundError).JSON(c)
		return
	} else if status == database.DbFailed {
		result.Error(exception.UpdatePasswordError).JSON(c)
		return
	}
	_ = a.TokenService.DeleteAll(user.ID)

	result.Ok().JSON(c)
}
