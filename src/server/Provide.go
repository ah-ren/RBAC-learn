package server

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/database"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/profile"
	"github.com/Aoi-hosizora/RBAC-learn/src/service"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/sirupsen/logrus"
)

func provideServices(config *config.ServerConfig, logger *logrus.Logger) *xdi.DiContainer {
	dic := xdi.NewDiContainer()

	dic.Provide(config)
	dic.Provide(logger)

	dic.Provide(profile.CreateEntityMappers())
	dic.Provide(database.SetupMySqlConn(config.MySqlConfig, logger))

	dic.Provide(service.NewUserRepository(dic))
	dic.Provide(service.NewJwtService(dic))
	dic.Provide(service.NewCasbinService(dic))

	return dic
}
