package usecase

import (
	"fmt"

	"github.com/hb-chen/micro-starter/service/greeting/domain/model"
	"github.com/hb-chen/micro-starter/service/greeting/domain/repository"
)

type GreetingUsecase interface {
	Add(msg string) (*model.Msg, error)
	List(page, size int) ([]*model.Msg, error)
	Duplicated(msg string) error
}

type greetingUsecase struct {
	repo repository.GreetingRepository
}

func NewGreetingUsecase(repo repository.GreetingRepository) GreetingUsecase {
	return &greetingUsecase{
		repo: repo,
	}
}

func (s *greetingUsecase) Add(msg string) (*model.Msg, error) {
	item, err := s.repo.FindByMsg(msg)
	if err != nil {
		return nil, err
	}

	if item != nil {
		return item, nil
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

func (s *greetingUsecase) List(page, size int) ([]*model.Msg, error) {
	items, err := s.repo.List(page, size)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *greetingUsecase) Duplicated(msg string) error {
	item, err := s.repo.FindByMsg(msg)
	if item != nil {
		return fmt.Errorf("msg duplicated: %s", item.Msg)
	}
	if err != nil {
		return err
	}
	return nil
}
