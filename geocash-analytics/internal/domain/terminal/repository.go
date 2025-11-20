// internal/domain/terminal/repository.go

package terminal

import (
	"context"
)

// Repository определяет базовые операции по работе с сущностями "Terminal".
type Repository interface {
	// GetByID извлекает полную информацию об устройстве по ID.
	GetByID(ctx context.Context, id int) (Terminal, error)

	// GetAllActive возвращает список всех активных устройств.
	GetAllActive(ctx context.Context) ([]Terminal, error)

	// ... Add, Update, Delete методы, если они необходимы.
}
