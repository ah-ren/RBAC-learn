package util

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/common/exception"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type authUtil struct{}

var AuthUtil = &authUtil{}

type Claims struct {
	jwt.StandardClaims
	UserId uint32 `json:"user-id"`
}

func (a *authUtil) EncryptPassword(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pwd), nil
}

func (a *authUtil) CheckPassword(password string, encrypted string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password))
	if err == nil {
		return true, nil
	} else if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	} else {
		return false, err
	}
}

func (a *authUtil) signToken(secret string, claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (a *authUtil) GenerateToken(uid uint32, isRefresh bool, config *config.JwtConfig) (string, error) {
	ex := config.Expire
	if isRefresh {
		ex = config.RefreshExpire
	}
	claims := &Claims{
		UserId: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + ex,
			Issuer:    config.Issuer,
		},
	}
	return a.signToken(config.Secret, claims)
}

func (a *authUtil) ParseToken(accessToken string, config *config.JwtConfig) (*Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	}
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ValidationError{Errors: jwt.ValidationErrorClaimsInvalid}
	}
	return claims, nil
}

func (a *authUtil) RefreshToken(refreshToken string, accessToken string, config *config.JwtConfig) (refresh string, access string, err error) {
	refreshClaims, err := a.ParseToken(refreshToken, config)
	if err != nil {
		if a.IsTokenExpired(err) {
			return "", "", exception.InvalidRefreshTokenError
		}
		return "", "", err
	}

	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	accessClaims, err := a.ParseToken(accessToken, config)
	if err != nil {
		return "", "", err
	}
	jwt.TimeFunc = time.Now

	if refreshClaims.UserId != accessClaims.UserId || refreshClaims.Issuer != accessClaims.Issuer {
		return "", "", exception.InvalidRefreshTokenError
	}
	newRefreshToken, err := a.GenerateToken(refreshClaims.UserId, true, config)
	newAccessToken, err1 := a.GenerateToken(refreshClaims.UserId, false, config)
	if err == nil {
		err = err1
	}
	return newRefreshToken, newAccessToken, err
}

func (a *authUtil) IsTokenExpired(err error) bool {
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorExpired != 0 {
			return true
		}
	}
	return false
}
