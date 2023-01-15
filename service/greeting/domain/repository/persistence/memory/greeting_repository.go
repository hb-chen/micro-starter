package memory

import (
	"sync"

	"github.com/hb-chen/micro-starter/service/greeting/domain/model"
	"github.com/hb-chen/micro-starter/service/greeting/domain/repository"
)

type greetingRepository struct {
	mu    *sync.Mutex
	items []*model.Msg
}

func (r *greetingRepository) FindById(id int64) (*model.Msg, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, item := range r.items {
		if item.Id == id {
			return item, nil
		}
	}
	return nil, nil
}

func (r *greetingRepository) FindByMsg(msg string) (*model.Msg, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, item := range r.items {
		if item.Msg == msg {
			return item, nil
		}
	}
	return nil, nil
}

func (r *greetingRepository) Add(item *model.Msg) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := int64(len(r.items) + 1)
	item.Id = id

	r.items = append(r.items, item)

	return nil
}

func (r *greetingRepository) List(page, size int) ([]*model.Msg, error) {
	return nil, nil
}

func NewGreetingRepository() repository.GreetingRepository {
	items := make([]*model.Msg, 0)

	return &greetingRepository{
		mu:    &sync.Mutex{},
		items: items,
	}
}
