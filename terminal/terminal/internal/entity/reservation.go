package entity

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TerminalID  uuid.UUID
	UserID      string
	Amount      int64
	Status      string
	ReservedAt  time.Time
	CompletedAt *time.Time
}
