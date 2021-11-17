package service

import (
	"fmt"

	"github.com/hb-chen/micro-starter/service/account/domain/model"
	"github.com/hb-chen/micro-starter/service/account/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Login(name, pwd string) (*model.User, error) {
	user, err := s.repo.FindByName(name)

	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, nil
	}

	if user.Password != pwd {
		return nil, nil
	}

	return user, nil

}

func (s *UserService) Register(name, pwd string) (*model.User, error) {
	err := s.Duplicated(name)
	if err != nil {
		return nil, err
	}

	u := model.User{
		Name:     name,
		Password: pwd,
	}
	err = s.repo.Add(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *UserService) Duplicated(name string) error {
	user, err := s.repo.FindByName(name)
	if user != nil {
		return fmt.Errorf("%s already exists", name)
	}
	if err != nil {
		return err
	}
	return nil
}
