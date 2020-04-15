package service

import (
	"fmt"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type TokenService struct {
	Config *config.ServerConfig `di:"~"`
	Conn   redis.Conn           `di:"~"`
}

func NewTokenService(dic *xdi.DiContainer) *TokenService {
	srv := &TokenService{}
	dic.InjectForce(srv)
	return srv
}

func (t *TokenService) concat(uid string, token string) string {
	return fmt.Sprintf(t.Config.JwtConfig.RedisFmt, uid, token)
}

func (t *TokenService) Query(token string) bool { // auth
	pattern := t.concat("*", token)
	keys, err := redis.Strings(t.Conn.Do("KEYS", pattern))
	if err != nil {
		return false
	}
	return len(keys) >= 1
}

func (t *TokenService) Insert(token string, uid uint32) bool { // login
	value := t.concat(strconv.Itoa(int(uid)), token)
	_, err := t.Conn.Do("SET", value, uid, "EX", t.Config.JwtConfig.Expire)
	return err == nil
}

func (t *TokenService) Delete(token string) bool { // logout
	pattern := t.concat("*", token)
	return t.deleteAll(pattern)
}

func (t *TokenService) DeleteAll(uid uint32) bool { // changePass
	pattern := t.concat(strconv.Itoa(int(uid)), "*")
	return t.deleteAll(pattern)
}

func (t *TokenService) deleteAll(pattern string) bool {
	keys, err := redis.Strings(t.Conn.Do("KEYS", pattern))
	if err != nil {
		return false
	}

	cnt := 0
	for _, key := range keys {
		result, err := redis.Int(t.Conn.Do("DEL", key))
		if err == nil {
			cnt += result
		}
	}
	return len(keys) == 0 || cnt > 0
}
