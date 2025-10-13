package service

import (
	"context"
	"tedx-maps/internal/entity"
	"tedx-maps/pkg/mapsapi"
)

type MapService struct {
	client mapsapi.Client
}

func NewMapService(client mapsapi.Client) *MapService {
	return &MapService{client: client}
}

func (s *MapService) GetMapData(ctx context.Context, query string) (*entity.MapData, error) {
	data, err := s.client.FetchMapData(ctx, query)
	if err != nil {
		return nil, err
	}
	return &entity.MapData{
		Source: "OSM",
		Data:   data,
	}, nil
}
