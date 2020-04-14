package util

import (
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

func (a *authUtil) GenerateToken(uid uint32, config *config.JwtConfig) (string, error) {
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
	if !ok || !token.Valid {
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
	if !ok || !token.Valid {
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
