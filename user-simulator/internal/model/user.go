package model

import (
	"errors"
	"math/rand"
	"time"
)

type User struct {
	ID             int       `json:"id"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	Balance        float64   `json:"balance"`
	ReservedAmount float64   `json:"reserved_amount"`
	LastUpdated    time.Time `json:"last_updated"`
}

func NewUser(id int, balance float64, lat, lon float64) *User {
	return &User{
		ID:          id,
		Balance:     balance,
		Latitude:    lat,
		Longitude:   lon,
		LastUpdated: time.Now(),
	}
}

func (u *User) Reserve(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть > 0")
	}
	if amount > u.Balance {
		return errors.New("недостаточно средств")
	}
	u.ReservedAmount = amount
	return nil
}

func (u *User) MoveRandom() {
	u.Latitude += (rand.Float64() - 0.5) / 1000
	u.Longitude += (rand.Float64() - 0.5) / 1000
	u.LastUpdated = time.Now()
}
