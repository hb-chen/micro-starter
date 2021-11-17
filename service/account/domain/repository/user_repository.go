package repository

import (
	"github.com/hb-chen/micro-starter/service/account/domain/model"
)

type UserRepository interface {
	FindById(id int64) (*model.User, error)
	FindByName(name string) (*model.User, error)
	Add(user *model.User) error
	List(page, size int) ([]*model.User, error)
}
