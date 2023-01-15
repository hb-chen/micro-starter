package model

type Msg struct {
	Id  int64  `xorm:"bigint pk autoincr" gorm:"primary_key"`
	Msg string `xorm:"varchar(128) unique" gorm:"type:varchar(128);unique_index"`
}
