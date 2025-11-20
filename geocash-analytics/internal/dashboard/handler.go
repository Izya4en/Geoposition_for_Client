package dashboard

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"geocash/internal/analytics" // Импорт бизнес-логики
)

// Handler группирует методы обработки запросов
type Handler struct {
	analyticsService analytics.Service
}

// NewHandler создает новый экземпляр хэндлера с внедренными зависимостями
func NewHandler(service analytics.Service) *Handler {
	return &Handler{
		analyticsService: service,
	}
}

// GetTerminalEfficiency обрабатывает GET /api/v1/terminals/{id}/efficiency
func (h *Handler) GetTerminalEfficiency(w http.ResponseWriter, r *http.Request) {
	// 1. Получаем ID из URL (предполагаем, что роутер передал его в query params или path)
	// Для примера берем из query string: ?id=1
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, `{"error": "missing id parameter"}`, http.StatusBadRequest)
		return
	}

	terminalID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "invalid id parameter"}`, http.StatusBadRequest)
		return
	}

	// 2. Вызываем бизнес-логику
	// Передаем контекст запроса, чтобы при отмене запроса клиентом отменялась и операция в БД
	status, err := h.analyticsService.CalculateEfficiencyStatus(r.Context(), terminalID)
	if err != nil {
		// В реальном проекте здесь нужно различать ошибки (404 vs 500)
		// Для простоты вернем 500
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	// 3. Формируем успешный ответ
	response := EfficiencyResponse{
		TerminalID: terminalID,
		Status:     status,
		CheckedAt:  time.Now(),
	}

	// 4. Отправляем JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
