package repository

import (
	"context"
	"database/sql"
	"tedx-maps/internal/entity"
)

type routeRepo struct {
	db *sql.DB
}

func NewRouteRepository(db *sql.DB) RouteRepository {
	return &routeRepo{db: db}
}

func (r *routeRepo) Create(ctx context.Context, route *entity.Route) error {
	query := `INSERT INTO routes (from_point, to_point, distance_km, duration, path, created_at)
			  VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id`
	return r.db.QueryRowContext(ctx, query,
		route.FromPoint, route.ToPoint, route.DistanceKM, route.Duration, route.Path).Scan(&route.ID)
}

func (r *routeRepo) GetByID(ctx context.Context, id int64) (*entity.Route, error) {
	query := `SELECT id, from_point, to_point, distance_km, duration, path, created_at FROM routes WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	var rt entity.Route
	err := row.Scan(&rt.ID, &rt.FromPoint, &rt.ToPoint, &rt.DistanceKM, &rt.Duration, &rt.Path, &rt.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &rt, nil
}

func (r *routeRepo) GetAll(ctx context.Context) ([]entity.Route, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, from_point, to_point, distance_km, duration, path, created_at FROM routes`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []entity.Route
	for rows.Next() {
		var rt entity.Route
		if err := rows.Scan(&rt.ID, &rt.FromPoint, &rt.ToPoint, &rt.DistanceKM, &rt.Duration, &rt.Path, &rt.CreatedAt); err != nil {
			return nil, err
		}
		routes = append(routes, rt)
	}
	return routes, nil
}

func (r *routeRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM routes WHERE id=$1`, id)
	return err
}
