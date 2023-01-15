package usecase

import (
	"github.com/hb-chen/micro-starter/service/greeting/domain/repository"
	"github.com/hb-chen/micro-starter/service/greeting/domain/service"
	"github.com/hb-go/pkg/conv"
)

type Msg struct {
	Id  int64  `json:"id"`
	Msg string `json:"msg"`
}

type GreetingUseCase interface {
	Call(msg string) (*Msg, error)
	List(page, size int) ([]*Msg, error)
}

type greetingUseCase struct {
	repo    repository.GreetingRepository
	service *service.GreetingService
}

func (uc *greetingUseCase) Call(msg string) (*Msg, error) {
	result, err := uc.service.Add(msg)
	if err != nil {
		return nil, err
	}

	item := &Msg{}
	conv.StructToStruct(result, item)

	return item, nil
}

func (uc *greetingUseCase) List(page, size int) ([]*Msg, error) {
	list, err := uc.repo.List(page, size)
	if err != nil {
		return nil, err
	}

	items := make([]*Msg, 0, len(list))
	for _, m := range list {
		item := &Msg{}
		conv.StructToStruct(m, item)
		items = append(items, item)
	}

	return items, nil
}

func NewGreetingUseCase(repo repository.GreetingRepository, service *service.GreetingService) GreetingUseCase {
	return &greetingUseCase{
		repo:    repo,
		service: service,
	}
}
