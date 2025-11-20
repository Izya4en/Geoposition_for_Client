// internal/analytics/service.go

package analytics

import (
	"context"
	"fmt"
	"time"
	// Для форматирования сообщения о статусе
)

// Service определяет интерфейс для бизнес-логики аналитики.
type Service interface {
	// CalculateEfficiencyStatus определяет общий статус устройства (эффективно/неэффективно).
	CalculateEfficiencyStatus(ctx context.Context, terminalID int) (string, error)
}

// AnalyticsService содержит зависимости сервиса (например, репозиторий).
type AnalyticsService struct {
	repo Repository // Наш интерфейс репозитория
	// Здесь также может быть Logger, Config и другие зависимости
}

// NewService создает и возвращает новый экземпляр AnalyticsService.
func NewService(repo Repository) *AnalyticsService {
	return &AnalyticsService{
		repo: repo,
	}
}

const (
	// Пороговые значения для определения эффективности
	MinDailyTransactions = 10.0        // Минимум 10 транзакций в день
	MinThroughputKZT     = 5_000_000.0 // Минимум 5 миллионов KZT оборота в неделю
	MinLoadingPercent    = 0.15        // Минимум 15% средней загрузки (чтобы не было простоя)
	MaxLoadingPercent    = 0.85        // Максимум 85% средней загрузки (чтобы не было частых дозагрузок)
)

// CalculateEfficiencyStatus реализует бизнес-логику оценки эффективности.
func (s *AnalyticsService) CalculateEfficiencyStatus(ctx context.Context, terminalID int) (string, error) {
	// 1. Определяем период для анализа (например, последние 7 дней)
	end := time.Now()
	start := end.AddDate(0, 0, -7)

	// 2. Получаем агрегированные метрики
	metrics, err := s.repo.GetPerformanceMetricsByPeriod(ctx, terminalID, start, end)
	if err != nil {
		return "", fmt.Errorf("ошибка при получении метрик: %w", err)
	}

	// 3. Получаем текущий остаток для оценки немедленной загруженности
	balance, err := s.repo.GetLastKnownBalance(ctx, terminalID)
	if err != nil {
		// Логируем ошибку, но продолжаем, если это не критично
		// return "", fmt.Errorf("ошибка при получении остатка: %w", err)
	}

	// --- 4. Применяем Правила Оценки ---

	// Переменная для сбора факторов неэффективности
	reasons := make([]string, 0)
	isEffective := true

	// 4.1. Проверка Проходимости (Транзакции)
	dailyTransactions := float64(metrics.TotalTransactions) / 7.0
	if dailyTransactions < MinDailyTransactions {
		isEffective = false
		reasons = append(reasons, fmt.Sprintf("Низкая проходимость (%.1f транз/день < %.1f)", dailyTransactions, MinDailyTransactions))
	}

	// 4.2. Проверка Оборота (Сумма)
	if metrics.TotalThroughputAmount < MinThroughputKZT {
		isEffective = false
		reasons = append(reasons, fmt.Sprintf("Низкий оборот (%.2f KZT/нед < %.2f)", metrics.TotalThroughputAmount, MinThroughputKZT))
	}

	// 4.3. Проверка Загруженности (Средняя)
	if metrics.AverageLoadingPercent < MinLoadingPercent {
		isEffective = false
		reasons = append(reasons, fmt.Sprintf("Низкая средняя загрузка (%.1f%%), простой", metrics.AverageLoadingPercent*100))
	}
	if metrics.AverageLoadingPercent > MaxLoadingPercent {
		// Это может быть неэффективно из-за частых инкассаций
		reasons = append(reasons, fmt.Sprintf("Высокая средняя загрузка (%.1f%%), частая инкассация", metrics.AverageLoadingPercent*100))
	}

	// 4.4. Проверка Истории Обслуживания (Критичность)
	if metrics.LastServiceCriticality {
		isEffective = false
		reasons = append(reasons, "Был критический ремонт в течение 7 дней")
	}

	// 4.5. Проверка Текущего Остатка (Прогноз)
	currentLoadingPercent := balance.CurrentBalance / balance.MaxCapacity
	if currentLoadingPercent < 0.05 {
		reasons = append(reasons, "Критически низкий текущий остаток (скоро будет пуст)")
		// Этот фактор может быть не ключевым, но важным для мониторинга
	}

	// --- 5. Формирование Конечного Статуса ---
	if isEffective && len(reasons) == 0 {
		return "ЭФФЕКТИВНО", nil
	}

	// Если неэффективно, возвращаем статус и причины
	status := "НЕЭФФЕКТИВНО"
	if len(reasons) > 0 {
		status += fmt.Sprintf(" (Причины: %s)", reasons)
	}

	return status, nil
}
