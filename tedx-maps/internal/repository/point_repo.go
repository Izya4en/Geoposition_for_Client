package repository

import (
	"context"
	"database/sql"
	"tedx-maps/internal/entity"
)

type pointRepo struct {
	db *sql.DB
}

func NewPointRepository(db *sql.DB) PointRepository {
	return &pointRepo{db: db}
}

func (r *pointRepo) Create(ctx context.Context, p *entity.Point) error {
	query := `INSERT INTO points (name, latitude, longitude, type, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id`
	return r.db.QueryRowContext(ctx, query, p.Name, p.Latitude, p.Longitude, p.Type).Scan(&p.ID)
}

func (r *pointRepo) GetByID(ctx context.Context, id int64) (*entity.Point, error) {
	query := `SELECT id, name, latitude, longitude, type, created_at, updated_at FROM points WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	var p entity.Point
	err := row.Scan(&p.ID, &p.Name, &p.Latitude, &p.Longitude, &p.Type, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *pointRepo) GetAll(ctx context.Context) ([]entity.Point, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name, latitude, longitude, type, created_at, updated_at FROM points`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var points []entity.Point
	for rows.Next() {
		var p entity.Point
		if err := rows.Scan(&p.ID, &p.Name, &p.Latitude, &p.Longitude, &p.Type, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		points = append(points, p)
	}
	return points, nil
}

func (r *pointRepo) Update(ctx context.Context, p *entity.Point) error {
	_, err := r.db.ExecContext(ctx, `UPDATE points SET name=$1, latitude=$2, longitude=$3, type=$4, updated_at=NOW() WHERE id=$5`,
		p.Name, p.Latitude, p.Longitude, p.Type, p.ID)
	return err
}

func (r *pointRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM points WHERE id=$1`, id)
	return err
}
