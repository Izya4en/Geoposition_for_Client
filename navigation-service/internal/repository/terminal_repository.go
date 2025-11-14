package repository

import "navigation-service/internal/entity"

type TerminalRepository struct{}

func NewTerminalRepository() *TerminalRepository {
	return &TerminalRepository{}
}

func (r *TerminalRepository) GetAll() []entity.Terminal {
	return []entity.Terminal{
		{"T1", 43.238949, 76.889709, 50000},
		{"T2", 43.25667, 76.92861, 15000},
		{"T3", 43.27073, 76.94565, 20000},
	}
}
