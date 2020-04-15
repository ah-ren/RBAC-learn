package main

import (
	"flag"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/server"
	"log"
)

var (
	help       = *flag.Bool("h", false, "show help")
	configPath = *flag.String("config", "./src/config/config.yaml", "set config path")
)

func main() {
	flag.Parse()
	if help {
		flag.Usage()
	} else {
		run()
	}
}

// @Title            rbac learn
// @Version          1.0
// @Description      my rbac learning repository
// @Host             127.0.0.1:10001
// @BasePath         /
// @GlobalSecurity   Jwt Authorization header
func run() {
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to load config file: %v\n", err)
	}

	s := server.NewServer(cfg)
	s.Serve()
}
