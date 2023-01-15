package handler

import (
	"context"

	pb "github.com/hb-chen/micro-starter/service/greeting/proto/greeting"
	"github.com/hb-chen/micro-starter/service/greeting/usecase"
)

var _ pb.GreetingHandler = &Greeting{}

type Greeting struct {
	greetingUseCase usecase.GreetingUseCase
}

func (g Greeting) Call(ctx context.Context, req *pb.CallRequest, resp *pb.CallResponse) error {
	msg, err := g.greetingUseCase.Call(req.Msg)
	if err != nil {
		return err
	}

	resp.Id = msg.Id
	resp.Msg = msg.Msg

	return nil
}

func (g *Greeting) Call1(ctx context.Context, req *pb.CallRequest) (*pb.CallResponse, error) {
	msg, err := g.greetingUseCase.Call(req.Msg)
	if err != nil {
		return nil, err
	}

	return &pb.CallResponse{
		Id:  msg.Id,
		Msg: msg.Msg,
	}, nil
}

func NewGreetingService(useCase usecase.GreetingUseCase) *Greeting {
	return &Greeting{
		greetingUseCase: useCase,
	}
}
