package gorm

import (
	"time"

	"github.com/hb-chen/micro-starter/service/greeting/domain/model"
	"github.com/micro/micro/v3/service/config"
	log "github.com/micro/micro/v3/service/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"github.com/hb-chen/micro-starter/service/greeting/conf"
)

func NewDB() (*gorm.DB, error) {
	dbConf := conf.Database{}
	cv, err := config.Get("database")
	if err != nil {
		log.Fatal(err)
	}
	if err := cv.Scan(&dbConf); err != nil {
		log.Fatal(err)
	}

	dsn := dbConf.User + ":" + dbConf.Password + "@tcp(" + dbConf.Host + ":" + dbConf.Port + ")/" + dbConf.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = db.Use(
		dbresolver.Register(dbresolver.Config{}).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(dbConf.ConnMaxLifetime).
			SetMaxIdleConns(dbConf.MaxIdleConns).
			SetMaxOpenConns(dbConf.MaxOpenConns),
	)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return db, nil
}

func Migrate() error {
	db, err := NewDB()
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.Msg{})
	if err != nil {
		return err
	}

	return nil
}
