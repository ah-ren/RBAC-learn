package server

import (
	"fmt"
	_ "github.com/Aoi-hosizora/RBAC-learn/docs"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/middleware"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
)

type Server struct {
	Server *gin.Engine
	Config *config.ServerConfig
}

func NewServer(config *config.ServerConfig) *Server {
	engine := gin.New()
	gin.SetMode(config.MetaConfig.RunMode)

	setupBinding()
	logger := setupLogger(config)
	dic := provideServices(config, logger)

	engine.Use(gin.Recovery())
	engine.Use(middleware.CorsMiddleware())
	engine.Use(middleware.LoggerMiddleware(logger))

	if gin.Mode() == "debug" {
		ginpprof.Wrap(engine)
	}
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	setupApiRoute(engine, dic)
	setupCommonRoute(engine)

	return &Server{
		Server: engine,
		Config: config,
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
