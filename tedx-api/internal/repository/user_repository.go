package repository

import "tedx-api/internal/entity"

type UserRepository interface {
	GetByUsername(username string) (*entity.User, error)
}

type InMemoryUserRepo struct {
	users []entity.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: []entity.User{
			{ID: 1, Username: "alice", Password: "pass123", Role: entity.RoleUser},
			{ID: 2, Username: "bob", Password: "pass123", Role: entity.RoleAdmin},
		},
	}
}

func (r *InMemoryUserRepo) GetByUsername(username string) (*entity.User, error) {
	for _, u := range r.users {
		if u.Username == username {
			return &u, nil
		}
	}
	return nil, nil
}
