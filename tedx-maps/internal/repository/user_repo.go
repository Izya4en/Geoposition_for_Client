package repository

import (
	"context"
	"database/sql"
	"tedx-maps/internal/entity"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, u *entity.User) error {
	query := `INSERT INTO users (email, password, role, created_at)
			  VALUES ($1, $2, $3, NOW()) RETURNING id`
	return r.db.QueryRowContext(ctx, query, u.Email, u.Password, u.Role).Scan(&u.ID)
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, email, password, role, created_at FROM users WHERE email=$1`
	row := r.db.QueryRowContext(ctx, query, email)
	var u entity.User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	query := `SELECT id, email, password, role, created_at FROM users WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	var u entity.User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
