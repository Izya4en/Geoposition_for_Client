// internal/analytics/repository.go

package analytics

import (
	"context"
	"time"

	"geocash/internal/domain/terminal"
)

// PerformanceMetrics содержит ключевые показатели для расчета эффективности.
type PerformanceMetrics struct {
	TotalTransactions      int     // Количество операций
	TotalThroughputAmount  float64 // Общий оборот (сумма транзакций)
	AverageLoadingPercent  float64 // Средний процент загрузки
	LastServiceCriticality bool    // Были ли критические ремонты в недавнем прошлом
}

// Repository определяет методы для получения аналитических данных из хранилища.
type Repository interface {
	// GetPerformanceMetricsByPeriod извлекает агрегированные данные для оценки эффективности
	// за заданный период (например, последние 7 или 30 дней).
	GetPerformanceMetricsByPeriod(
		ctx context.Context,
		terminalID int,
		start time.Time,
		end time.Time,
	) (PerformanceMetrics, error)

	// GetLastKnownBalance извлекает последнюю зафиксированную информацию об остатках
	// для расчета текущей загруженности.
	GetLastKnownBalance(ctx context.Context, terminalID int) (terminal.CashBalance, error)
}
