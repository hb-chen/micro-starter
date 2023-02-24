package repository

import (
	"github.com/hb-chen/micro-starter/service/greeting/domain/model"
)

type GreetingRepository interface {
	FindById(id int64) (*model.Msg, error)
	FindByMsg(msg string) (*model.Msg, error)
	Add(msg *model.Msg) error
	List(page, size int) ([]*model.Msg, error)
}
