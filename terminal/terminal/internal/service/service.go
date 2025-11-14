package service

import (
	"errors"
	"time"

	"terminal/internal/entity"
	"terminal/internal/repository"

	"github.com/google/uuid"
)

type Service struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) ReserveMoney(terminalID string, user string, amount int64) (*entity.Reservation, error) {
	t, err := s.Repo.GetTerminal(terminalID)
	if err != nil {
		return nil, err
	}
	if t.CashAvailable < amount {
		return nil, errors.New("not enough cash")
	}
	t.CashAvailable -= amount
	s.Repo.UpdateTerminal(t)

	res := &entity.Reservation{
		TerminalID: uuid.MustParse(terminalID),
		UserID:     user,
		Amount:     amount,
		Status:     "reserved",
		ReservedAt: time.Now(),
	}
	if err := s.Repo.SaveReservation(res); err != nil {
		return nil, err
	}
	return res, nil
}
