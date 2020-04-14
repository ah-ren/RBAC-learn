package server

import (
	"fmt"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/ahlib-gin-gorm/xgin"
	"github.com/gin-gonic/gin"
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

func setupBinding() {
	xgin.SetupRegexBinding()
}

func setupLogger(config *config.Config) *logrus.Logger {
	logger := logrus.New()
	logLevel := logrus.WarnLevel
	if config.MetaConfig.RunMode == "debug" {
		logLevel = logrus.DebugLevel
	}

	// file
	fileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   config.MetaConfig.LogPath,
		MaxSize:    50,
		MaxBackups: 3,
		MaxAge:     30,
		Level:      logLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		},
	})
	if err != nil {
		log.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	// logrus
	logger.SetLevel(logLevel)
	logger.SetReportCaller(true)
	logger.AddHook(fileHook)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		DisableSorting:  true,
		TimestampFormat: "2006/01/02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf(" %s:%d:", filename, f.Line)
		},
	})

	// writer
	out := io.MultiWriter(ansicolor.NewAnsiColorWriter(os.Stdout))
	log.SetOutput(out)
	gin.DefaultWriter = out
	logger.Out = out

	return logger
}
