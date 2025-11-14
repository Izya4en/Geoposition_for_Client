package service

import (
	"log"
	"math/rand"
	"time"

	"user-simulator/internal/model"
	"user-simulator/internal/transport"
)

type Simulator struct {
	User   *model.User
	Client *transport.HTTPClient
}

func NewSimulator(user *model.User, client *transport.HTTPClient) *Simulator {
	return &Simulator{User: user, Client: client}
}

func (s *Simulator) Start(interval time.Duration) {
	for {
		s.User.MoveRandom()

		amount := rand.Float64() * 15000
		if err := s.User.Reserve(amount); err != nil {
			log.Printf("⚠️ Резервация не удалась: %v", err)
		}

		if err := s.Client.SendStatus(s.User); err != nil {
			log.Printf("❌ Ошибка отправки: %v", err)
		} else {
			log.Printf("📤 Данные отправлены: user=%d reserved=%.2f", s.User.ID, s.User.ReservedAmount)
		}

		time.Sleep(interval)
	}
}
