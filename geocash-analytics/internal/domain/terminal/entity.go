// internal/domain/terminal/entity.go

package terminal

import "time"

// CashBalance отражает текущее состояние наличности в устройстве
// (соответствует данным из таблицы terminal_loadings).
type CashBalance struct {
	TerminalID     int
	RecordTime     time.Time
	CurrentBalance float64
	MaxCapacity    float64
}

// Terminal Entity, если нужно получить полные данные об устройстве.
type Terminal struct {
	ID           int
	SerialNumber string
	Location     string
	Model        string
	IsActive     bool
	// ... другие поля из таблицы terminals
}
