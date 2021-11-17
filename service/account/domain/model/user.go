package model

type User struct {
	Id       int64  `xorm:"bigint pk autoincr" gorm:"primary_key"`
	Name     string `xorm:"varchar(128) unique" gorm:"type:varchar(128);unique_index"`
	Password string `xorm:"varchar(255)" gorm:"type:varchar(255)"`
}
