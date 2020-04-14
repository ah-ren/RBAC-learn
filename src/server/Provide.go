package server

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/database"
	"github.com/Aoi-hosizora/RBAC-learn/src/middleware"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/profile"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/sirupsen/logrus"
)

func provideServices(config *config.Config, logger *logrus.Logger) *xdi.DiContainer {
	dic := xdi.NewDiContainer()

	dic.Provide(config)
	dic.Provide(logger)

	dic.Provide(profile.CreateEntityMappers())
	dic.Provide(database.NewUserRepository())
	dic.Provide(database.NewCasbinRepository().Adapter())
	dic.Provide(middleware.NewJwtService(dic))
	dic.Provide(middleware.NewCasbinService(dic))

	return dic
}
