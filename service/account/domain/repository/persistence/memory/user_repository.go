package memory

import (
	"sync"

	"github.com/hb-chen/micro-starter/service/account/domain/model"
	"github.com/hb-chen/micro-starter/service/account/domain/repository"
)

type userRepository struct {
	mu    *sync.Mutex
	users []*model.User
}

func NewUserRepository() repository.UserRepository {
	users := make([]*model.User, 0)
	users = append(users, &model.User{
		Id:       1,
		Name:     "admin",
		Password: "123456",
	})

	return &userRepository{
		mu:    &sync.Mutex{},
		users: users,
	}
}

func (r *userRepository) FindById(id int64) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, nil
}

func (r *userRepository) FindByName(name string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Name == name {
			return user, nil
		}
	}
	return nil, nil
}

func (r *userRepository) Add(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := int64(len(r.users) + 1)
	user.Id = id

	r.users = append(r.users, user)

	return nil
}

func (r *userRepository) List(page, size int) ([]*model.User, error) {
	return nil, nil
}
