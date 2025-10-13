package service

import (
	"context"
	"errors"
	"tedx-maps/internal/entity"
	"tedx-maps/internal/repository"
)

type PointService struct {
	repo repository.PointRepository
}

func NewPointService(repo repository.PointRepository) *PointService {
	return &PointService{repo: repo}
}

func (s *PointService) CreatePoint(ctx context.Context, point *entity.Point) (*entity.Point, error) {
	if point == nil || point.Latitude == 0 || point.Longitude == 0 {
		return nil, errors.New("invalid point data")
	}
	if err := s.repo.Create(ctx, point); err != nil {
		return nil, err
	}
	return point, nil
}

func (s *PointService) GetAllPoints(ctx context.Context) ([]entity.Point, error) {
	return s.repo.GetAll(ctx)
}

func (s *PointService) GetNearestPoints(ctx context.Context, lat, lon float64, radius float64) ([]entity.Point, error) {
	return s.repo.GetNearest(ctx, lat, lon, radius)
}
