package dashboard

import "time"

// EfficiencyResponse — структура ответа API
type EfficiencyResponse struct {
	TerminalID int       `json:"terminal_id"`
	Status     string    `json:"status"`     // "ЭФФЕКТИВНО" / "НЕЭФФЕКТИВНО"
	CheckedAt  time.Time `json:"checked_at"` // Время проверки
}

// ErrorResponse — стандартная ошибка
type ErrorResponse struct {
	Error string `json:"error"`
}
