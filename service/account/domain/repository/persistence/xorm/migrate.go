package xorm

import (
	"github.com/hb-chen/micro-starter/service/account/domain/model"
	"xorm.io/xorm"
	"xorm.io/xorm/migrate"
)

var (
	migrations = []*migrate.Migration{
		{
			ID: "201911170000",
			Migrate: func(tx *xorm.Engine) error {
				// user表
				if err := tx.Sync2(&model.User{}); err != nil {
					return err
				}

				// 默认管理员
				admin := &model.User{
					Name:     "admin",
					Password: "123456",
				}
				if _, err := tx.Insert(admin); err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *xorm.Engine) error {
				return tx.DropTables(&model.User{})
			},
		},
	}
)
