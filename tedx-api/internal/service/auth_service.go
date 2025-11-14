package service

import (
	"errors"
	"time"

	"tedx-api/internal/config"
	"tedx-api/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	repo      repository.UserRepository
	jwtSecret string
}

func NewAuthService(repo repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{repo: repo, jwtSecret: cfg.JWTSecret}
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, _ := s.repo.GetByUsername(username)
	if user == nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"sub":  user.ID,
		"user": user.Username,
		"role": user.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
