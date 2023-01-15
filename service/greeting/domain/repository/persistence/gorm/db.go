package gorm

import (
	"sync"
	"time"

	"github.com/micro/micro/v3/service/config"
	log "github.com/micro/micro/v3/service/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"github.com/hb-chen/micro-starter/service/greeting/conf"
	"github.com/hb-chen/micro-starter/service/greeting/domain/model"
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

		db, err = gorm.Open(mysql.Open(sqlConnection), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		err = db.Use(
			dbresolver.Register(dbresolver.Config{}).
				SetConnMaxIdleTime(time.Hour).
				SetConnMaxLifetime(dbConf.ConnMaxLifetime).
				SetMaxIdleConns(dbConf.MaxIdleConns).
				SetMaxOpenConns(dbConf.MaxOpenConns),
		)
		if err != nil {
			log.Fatal(err)
			return
		}

		err = db.AutoMigrate(&model.Msg{})
		if err != nil {
			log.Fatal(err)
		}

		// TODO 数据初始化
	})
}
