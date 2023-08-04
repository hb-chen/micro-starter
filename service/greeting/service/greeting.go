package service

import (
	"context"

	"github.com/hb-chen/micro-starter/service/greeting/domain/repository"
	"github.com/hb-chen/micro-starter/service/greeting/domain/usecase"
	"github.com/hb-chen/micro-starter/service/greeting/proto/greeting"
)

type Msg struct {
	Id  int64  `json:"id"`
	Msg string `json:"msg"`
}

var _ greeting.GreetingHandler = &greetingService{}

type greetingService struct {
	repo            repository.GreetingRepository
	greetingUsecase usecase.GreetingUsecase
}

func (g *greetingService) Call(ctx context.Context, req *greeting.CallRequest, resp *greeting.CallResponse) error {
	result, err := g.greetingUsecase.Add(req.Msg)
	if err != nil {
		return err
	}

	resp.Id = result.Id
	resp.Msg = result.Msg

	return nil
}

func (g *greetingService) List(ctx context.Context, page *greeting.Page, resp *greeting.ListResponse) error {
	items, err := g.greetingUsecase.List(int(page.Page), int(page.Size))
	if err != nil {
		return err
	}

	for _, item := range items {

		resp.Items = append(resp.Items,
			&greeting.CallResponse{
				Id:  item.Id,
				Msg: item.Msg,
			})
	}

	return nil
}

func NewGreetingService(repo repository.GreetingRepository, us usecase.GreetingUsecase) greeting.GreetingHandler {
	return &greetingService{
		repo:            repo,
		greetingUsecase: us,
	}
}
