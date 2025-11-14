package entity

import (
	"time"

	"github.com/google/uuid"
)

type Terminal struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name          string
	Latitude      float64
	Longitude     float64
	CashAvailable int64
	Status        string
	CreatedAt     time.Time
}
