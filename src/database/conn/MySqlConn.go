package conn

import (
	"fmt"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/po"
	"github.com/Aoi-hosizora/ahlib-gin-gorm/xgorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"log"
)

func SetupMySqlConn(config *config.MySqlConfig, logger *logrus.Logger) *gorm.DB {
	dbParams := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.User, config.Password,
		config.Host, config.Port,
		config.Name, config.Charset,
	)
	db, err := gorm.Open("mysql", dbParams)
	if err != nil {
		log.Fatalln("Failed to connect mysql:", err)
	}

	db.LogMode(config.IsLog)
	db.SetLogger(NewGormLogger(logger))

	xgorm.HookDeleteAtField(db, xgorm.DefaultDeleteAtTimeStamp)
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tbl_" + defaultTableName
	}

	autoMigrateModel(db)

	return db
}

func autoMigrateModel(db *gorm.DB) {
	autoMigrate := func(value interface{}) {
		rdb := db.AutoMigrate(value)
		if rdb.Error != nil {
			log.Fatalln(rdb.Error)
		}
	}

	autoMigrate(&po.User{})
}
