package server

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/database"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/profile"
	"github.com/Aoi-hosizora/RBAC-learn/src/service"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

func provideServices(config *config.ServerConfig, logger *logrus.Logger) *xdi.DiContainer {
	dic := xdi.NewDiContainer()

	dic.Provide(config)
	dic.Provide(logger)

	dic.Provide(profile.CreateEntityMappers())
	dic.Provide(database.SetupMySQLConn(config.MySQLConfig, logger))
	dic.ProvideImpl((*redis.Conn)(nil), database.SetupRedisConn(config.RedisConfig, logger))

	dic.Provide(service.NewUserService(dic))
	dic.Provide(service.NewTokenService(dic))
	dic.Provide(service.NewJwtService(dic))
	dic.Provide(service.NewCasbinService(dic))

	return dic
}
