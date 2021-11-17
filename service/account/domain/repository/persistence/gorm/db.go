package gorm

import (
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/micro/micro/v3/service/config"
	log "github.com/micro/micro/v3/service/logger"

	"github.com/hb-chen/micro-starter/service/account/conf"
	"github.com/hb-chen/micro-starter/service/account/domain/model"
)

var (
	dbConf conf.Database
	db     *gorm.DB
	once   sync.Once
)

func InitDB() {
	once.Do(func() {
		dbConf = conf.Database{}
		cv, err := config.Get("database")
		if err != nil {
			log.Fatal(err)
		}
		if err := cv.Scan(&dbConf); err != nil {
			log.Fatal(err)
		}

		sqlConnection := dbConf.User + ":" + dbConf.Password + "@tcp(" + dbConf.Host + ":" + dbConf.Port + ")/" + dbConf.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(dbConf.Engine, sqlConnection)
		if err != nil {
			log.Fatal(err)
		}

		db.DB().SetMaxOpenConns(dbConf.MaxOpenConns)
		db.DB().SetMaxIdleConns(dbConf.MaxIdleConns)
		db.DB().SetConnMaxLifetime(dbConf.ConnMaxLifetime)

		db.SingularTable(true)
		err = db.AutoMigrate(&model.User{}).Error
		if err != nil {
			log.Fatal(err)
		}

		// TODO 数据初始化
	})
}
