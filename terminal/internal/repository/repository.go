package repository

import (
	"gorm.io/gorm"
	"terminal/internal/entity"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateTerminal(t *entity.Terminal) error {
	return r.DB.Create(t).Error
}

func (r *Repository) ListTerminals() ([]entity.Terminal, error) {
	var terms []entity.Terminal
	return terms, r.DB.Find(&terms).Error
}

func (r *Repository) GetTerminal(id string) (*entity.Terminal, error) {
	var t entity.Terminal
	if err := r.DB.First(&t, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *Repository) SaveReservation(res *entity.Reservation) error {
	return r.DB.Create(res).Error
}

func (r *Repository) UpdateTerminal(t *entity.Terminal) error {
	return r.DB.Save(t).Error
}
