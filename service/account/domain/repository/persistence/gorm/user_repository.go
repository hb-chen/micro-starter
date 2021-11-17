package gorm

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
	if result := db.Where("id = ?", id).First(&user); result.Error != nil {
		if result.RecordNotFound() {
			return nil, nil
		} else {
			return nil, result.Error
		}
	} else {
		return &user, nil
	}
}

func (r *userRepository) FindByName(name string) (*model.User, error) {
	user := model.User{}
	if result := db.Where("name = ?", name).First(&user); result.Error != nil {
		if result.RecordNotFound() {
			return nil, nil
		} else {
			return nil, result.Error
		}
	} else {
		return &user, nil
	}
}

func (r *userRepository) Add(user *model.User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) List(page, size int) ([]*model.User, error) {
	list := make([]*model.User, 0)
	err := db.Order("id DESC").Offset((page - 1) * size).Limit(size).Find(&list).Error

	return list, err
}
