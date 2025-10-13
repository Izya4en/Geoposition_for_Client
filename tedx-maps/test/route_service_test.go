package test

import (
	"tedx-maps/internal/entity"
	"tedx-maps/internal/service"
	"testing"
)

func TestBuildRoute(t *testing.T) {
	svc := service.NewRouteService(nil)

	points := []entity.Point{
		{Latitude: 43.238949, Longitude: 76.889709},
		{Latitude: 43.240210, Longitude: 76.892101},
	}

	route := svc.BuildRoute(points)
	if route.DistanceKm == 0 {
		t.Errorf("expected non-zero route distance")
	}
}
