// +build ignore

package xorm

import (
	"github.com/hb-chen/micro-starter/service/account/domain/model"
	"github.com/hb-chen/micro-starter/service/account/domain/repository"
)

type userRepository struct {
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindById(id int64) (*model.User, error) {
	user := model.User{}
	if ok, err := db.Where("id = ?", id).Get(&user); ok && err == nil {
		return &user, nil
	} else {
		return nil, err
	}
}

func (r *userRepository) FindByName(name string) (*model.User, error) {
	user := model.User{}
	if has, err := db.Where("name = ?", name).Get(&user); err == nil && has {
		return &user, nil
	} else {
		return nil, err
	}
}

func (r *userRepository) Add(user *model.User) error {
	_, err := db.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) List(page, size int) ([]*model.User, error) {
	list := make([]*model.User, 0)
	session := db.Desc("id")
	err := session.Limit(size, (page-1)*size).Find(&list)

	return list, err
}
