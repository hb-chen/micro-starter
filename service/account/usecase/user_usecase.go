package usecase

import (
	"github.com/hb-chen/micro-starter/service/account/domain/repository"
	"github.com/hb-chen/micro-starter/service/account/domain/service"
	"github.com/hb-go/pkg/conv"
)

type UserUseCase interface {
	LoginUser(name, pwd string) (*User, error)
	RegisterUser(name, pwd string) (*User, error)
	GetUser(id int64) (*User, error)
	GetUserList(page, size int) ([]*User, error)
}

type userUseCase struct {
	repo    repository.UserRepository
	service *service.UserService
}

func NewUserUseCase(repo repository.UserRepository, service *service.UserService) UserUseCase {
	return &userUseCase{
		repo:    repo,
		service: service,
	}
}

func (uc *userUseCase) LoginUser(name, pwd string) (*User, error) {
	user, err := uc.service.Login(name, pwd)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, nil
	}

	u := &User{}
	conv.StructToStruct(user, u)
	return u, nil
}

func (uc *userUseCase) RegisterUser(name, pwd string) (*User, error) {
	user, err := uc.service.Register(name, pwd)
	if err != nil {
		return nil, err
	}

	u := &User{}
	conv.StructToStruct(user, u)
	return u, nil
}

func (uc *userUseCase) GetUser(id int64) (*User, error) {
	user, err := uc.repo.FindById(id)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, nil
	}

	return &User{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}

func (uc *userUseCase) GetUserList(page, size int) ([]*User, error) {
	list, err := uc.repo.List(page, size)
	if err != nil {
		return nil, err
	}

	users := make([]*User, 0, len(list))
	for _, u := range list {
		user := &User{}
		conv.StructToStruct(u, user)
		users = append(users, user)
	}

	return users, nil
}

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
