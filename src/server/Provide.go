package server

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/database/conn"
	"github.com/Aoi-hosizora/RBAC-learn/src/database/dao"
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
	dic.Provide(conn.SetupMySqlConn(config.MySqlConfig, logger))

	dic.Provide(dao.NewUserRepository(dic))

	dic.Provide(middleware.NewJwtService(dic))
	dic.Provide(middleware.NewCasbinService(dic))

	return dic
}
