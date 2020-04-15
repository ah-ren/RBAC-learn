package middleware

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/common/exception"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/database/dao"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/po"
	"github.com/Aoi-hosizora/RBAC-learn/src/util"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/gin-gonic/gin"
)

type JwtService struct {
	Config   *config.Config      `di:"~"`
	UserRepo *dao.UserRepository `di:"~"`

	_key string `di:"-"`
}

func NewJwtService(dic *xdi.DiContainer) *JwtService {
	srv := &JwtService{_key: "user"}
	dic.InjectForce(srv)
	return srv
}

func (j *JwtService) JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := j.getToken(c)
		user, err := j.jwtCheck(token)
		if err != nil {
			result.Error(err).JSON(c)
			c.Abort()
			return
		}

		c.Set(j._key, user)
		c.Next()
	}
}

func (j *JwtService) GetContextUser(c *gin.Context) *po.User {
	_user, exist := c.Get(j._key)
	if !exist {
		result.Error(exception.UnAuthorizedError).JSON(c)
		c.Abort()
		return nil
	}
	user, ok := _user.(*po.User)
	if !ok {
		result.Error(exception.UnAuthorizedError).JSON(c)
		c.Abort()
		return nil
	}
	return user
}

func (j *JwtService) getToken(c *gin.Context) string {
	if token := c.Request.Header.Get("Authorization"); token != "" {
		return token
	} else {
		return c.DefaultQuery("Authorization", "")
	}
}

func (j *JwtService) jwtCheck(token string) (*po.User, *exception.ServerError) {
	if token == "" {
		return nil, exception.UnAuthorizedError
	}

	claims, err := util.AuthUtil.ParseToken(token, j.Config.JwtConfig)
	if err != nil {
		if util.AuthUtil.IsTokenExpired(err) {
			return nil, exception.TokenExpiredError
		} else {
			return nil, exception.UnAuthorizedError
		}
	}

	user := j.UserRepo.QueryById(claims.UserId)
	if user == nil {
		return nil, exception.UnAuthorizedError
	}

	return user, nil
}
