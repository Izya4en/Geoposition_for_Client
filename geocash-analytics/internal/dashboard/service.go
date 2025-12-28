package dashboard

import (
	"fmt"
	"geocash/internal/analytics"
	"geocash/internal/domain/terminal"
	"geocash/internal/platform/provider"
	"strings"
)

type Service struct {
	repo terminal.Repository
	osm  *provider.OSMProvider
	grid *analytics.GridService

	forteCache []terminal.ATM
	compCache  []terminal.ATM
}

func NewService(repo terminal.Repository, osm *provider.OSMProvider, grid *analytics.GridService) *Service {
	s := &Service{repo: repo, osm: osm, grid: grid}
	go s.refreshData()
	return s
}

func (s *Service) refreshData() {
	fmt.Println("🔄 Updating ATM data from OpenStreetMap...")

	allATMs, err := s.osm.FetchAllATMs()
	if err != nil {
		fmt.Println("❌ OSM Error:", err)
		return
	}

	var forte []terminal.ATM
	var others []terminal.ATM

	for i := range allATMs {
		atm := allATMs[i]

		name := strings.ToLower(atm.Bank) + strings.ToLower(atm.Name)
		if strings.Contains(name, "forte") {

			s.repo.EnrichATM(&atm)
			forte = append(forte, atm)
		} else {

			atm.IsForte = false
			others = append(others, atm)
		}
	}

	s.forteCache = forte
	s.compCache = others
	fmt.Printf("✅ Data Updated: %d Forte ATMs, %d Competitors\n", len(forte), len(others))
}

func (s *Service) GetDashboardData() DashboardResponse {

	competitors := s.compCache
	if len(competitors) == 0 {
		competitors = s.repo.GenerateRandomCompetitors(300)
	}

	forte := s.forteCache

	return DashboardResponse{
		Forte:       forte,
		Competitors: competitors,
		HeatmapGrid: s.grid.GenerateHexGrid(),
	}
}
