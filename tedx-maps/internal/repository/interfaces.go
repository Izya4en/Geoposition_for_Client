package repository

import (
	"context"
	"tedx-maps/internal/entity"
)

type PointRepository interface {
	Create(ctx context.Context, p *entity.Point) error
	Update(ctx context.Context, p *entity.Point) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*entity.Point, error)
	GetAll(ctx context.Context) ([]entity.Point, error)
}

type RouteRepository interface {
	Create(ctx context.Context, r *entity.Route) error
	GetByID(ctx context.Context, id int64) (*entity.Route, error)
	GetAll(ctx context.Context) ([]entity.Route, error)
}

type UserRepository interface {
	Create(ctx context.Context, u *entity.User) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
}
