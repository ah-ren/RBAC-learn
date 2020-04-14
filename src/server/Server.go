package server

import (
	"fmt"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/middleware"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	Server *gin.Engine
	Config *config.Config
	Dic    *xdi.DiContainer
}

func NewServer(config *config.Config) *Server {
	engine := gin.New()
	gin.SetMode(config.MetaConfig.RunMode)
	if gin.Mode() == "debug" {
		ginpprof.Wrap(engine)
	}

	setupBinding()
	logger := setupLogger(config)
	dic := provideServices(config, logger)

	engine.Use(gin.Recovery())
	engine.Use(middleware.CorsMiddleware())
	engine.Use(middleware.LoggerMiddleware(logger))

	setupCommonRoute(engine)
	setupApiRoute(engine, dic)

	return &Server{
		Server: engine,
		Config: config,
		Dic:    dic,
	}
}

func (s *Server) Serve() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Config.MetaConfig.Port),
		Handler: s.Server,
	}

	fmt.Println()
	log.Printf("Server is listening on port %s\n\n", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to listen port and server: %v\n", err)
	}
}
