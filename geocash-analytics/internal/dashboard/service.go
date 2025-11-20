package dashboard

import (
	"geocash/internal/domain/terminal"
)

// Service оркестрирует получение данных
type Service struct {
	termRepo terminal.Repository
}
