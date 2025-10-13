package service

import (
	"context"
	"errors"
	"tedx-maps/internal/entity"
	"tedx-maps/internal/repository"
)

type RouteService struct {
	repo repository.RouteRepository
}

func NewRouteService(repo repository.RouteRepository) *RouteService {
	return &RouteService{repo: repo}
}

func (s *RouteService) CreateRoute(ctx context.Context, route *entity.Route) (*entity.Route, error) {
	if route == nil || route.StartPointID == 0 || route.EndPointID == 0 {
		return nil, errors.New("invalid route data")
	}
	if err := s.repo.Create(ctx, route); err != nil {
		return nil, err
	}
	return route, nil
}

func (s *RouteService) GetRoutes(ctx context.Context) ([]entity.Route, error) {
	return s.repo.GetAll(ctx)
}
