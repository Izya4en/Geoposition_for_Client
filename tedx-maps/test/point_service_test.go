package test

import (
	"testing"

	"tedx-maps/internal/entity"
	"tedx-maps/internal/service"
)

func TestCalculateDistance(t *testing.T) {
	svc := service.NewPointService(nil)

	p1 := entity.Point{Latitude: 43.238949, Longitude: 76.889709}
	p2 := entity.Point{Latitude: 43.240210, Longitude: 76.892101}

	distance := svc.CalculateDistance(p1, p2)
	if distance <= 0 {
		t.Errorf("expected positive distance, got %f", distance)
	}
}
