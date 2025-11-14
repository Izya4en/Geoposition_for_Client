package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Analytics struct {
	ID         uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TerminalID uuid.UUID
	EventType  string
	Payload    datatypes.JSON
	CreatedAt  time.Time
}
