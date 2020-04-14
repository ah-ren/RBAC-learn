package util

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type authUtil struct{}

var AuthUtil = &authUtil{}

type Claims struct {
	jwt.StandardClaims
	UserId uint32 `json:"user-id"`
}

func (a *authUtil) CreateToken(uid uint32, config *config.JwtConfig) (string, error) {
	claims := &Claims{
		UserId: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + config.Expire,
			Issuer:    config.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Secret))
}

func (a *authUtil) ParseToken(tokenString string, config *config.JwtConfig) (*Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || token.Valid {
		return nil, jwt.ValidationError{Errors: jwt.ValidationErrorClaimsInvalid}
	}
	return claims, nil
}

func (a *authUtil) RefreshToken(tokenString string, config *config.JwtConfig) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, keyFunc)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || token.Valid {
		return "", jwt.ValidationError{Errors: jwt.ValidationErrorClaimsInvalid}
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Unix() + config.Expire

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return newToken.SignedString([]byte(config.Secret))
}

func (a *authUtil) IsTokenExpired(err error) bool {
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorExpired != 0 {
			return true
		}
	}
	return false
}
