package service

import (
	"navigation-service/internal/entity"
	"navigation-service/internal/repository"
	"navigation-service/internal/utils"
)

type NavigationService struct {
	repo *repository.TerminalRepository
}

func NewNavigationService(repo *repository.TerminalRepository) *NavigationService {
	return &NavigationService{repo: repo}
}

func (s *NavigationService) FindNearest(lat, lon, amount float64) *entity.Terminal {
	terminals := s.repo.GetAll()
	var best *entity.Terminal
	minDist := 1e9

	for _, t := range terminals {
		if t.AvailableCash < amount {
			continue
		}
		dist := utils.CalculateDistance(lat, lon, t.Latitude, t.Longitude)
		if dist < minDist {
			minDist = dist
			best = &t
		}
	}
	return best
}
