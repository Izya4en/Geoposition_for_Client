package service

import (
	"context"
	"errors"
	"tedx-maps/internal/entity"
	"tedx-maps/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, user *entity.User) error {
	if user == nil || user.Email == "" || user.Password == "" {
		return errors.New("invalid user data")
	}
	return s.repo.Create(ctx, user)
}
