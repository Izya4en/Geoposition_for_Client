package service

import (
	"encoding/json"
	"reservation/internal/kafka"
	"reservation/internal/repository"
)

type ReservationService struct {
	repo     *repository.ReservationRepo
	producer *kafka.Producer
}

func NewReservationService(repo *repository.ReservationRepo, producer *kafka.Producer) *ReservationService {
	return &ReservationService{repo: repo, producer: producer}
}

func (s *ReservationService) ReserveMoney(userID string, amount float64) {
	reservation := repository.Reservation{
		UserID: userID,
		Amount: amount,
		Status: "reserved",
	}

	s.repo.Save(reservation)

	data, _ := json.Marshal(reservation)
	s.producer.SendMessage("reservations", string(data))
}

func (s *ReservationService) GetAll() []repository.Reservation {
	return s.repo.GetAll()
}
