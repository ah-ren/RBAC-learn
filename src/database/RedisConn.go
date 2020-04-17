package database

import (
	"fmt"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func SetupRedisConn(config *config.RedisConfig, logger *logrus.Logger) redis.Conn {
	conn, err := redis.Dial(
		config.ConnType,
		fmt.Sprintf("%s:%d", config.Host, config.Port),
		redis.DialPassword(config.Password),
		redis.DialDatabase(int(config.Db)),
		redis.DialConnectTimeout(time.Duration(config.ConnectTimeout)*time.Millisecond),
		redis.DialReadTimeout(time.Duration(config.ReadTimeout)*time.Millisecond),
		redis.DialWriteTimeout(time.Duration(config.WriteTimeout)*time.Millisecond),
	)
	if err != nil {
		log.Fatalln("Failed to connect redis:", err)
	}

	return NewRedisLogger(conn, logger)
}
