package service

import (
	"fmt"

	"github.com/hb-chen/micro-starter/service/greeting/domain/model"
	"github.com/hb-chen/micro-starter/service/greeting/domain/repository"
)

type GreetingService struct {
	repo repository.GreetingRepository
}

func NewGreetingService(repo repository.GreetingRepository) *GreetingService {
	return &GreetingService{
		repo: repo,
	}
}

func (s *GreetingService) Add(msg string) (*model.Msg, error) {
	err := s.Duplicated(msg)
	if err != nil {
		return nil, err
	}

	u := model.Msg{
		Msg: msg,
	}
	err = s.repo.Add(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *GreetingService) List(page, size int) ([]*model.Msg, error) {
	items, err := s.repo.List(page, size)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *GreetingService) Duplicated(msg string) error {
	item, err := s.repo.FindByMsg(msg)
	if item != nil {
		return fmt.Errorf("msg duplicated: %s", item.Msg)
	}
	if err != nil {
		return err
	}
	return nil
}
