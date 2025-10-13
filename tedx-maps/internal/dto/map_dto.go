package dto

import "tedx-maps/internal/entity"

// MapDataResponse — ответ API для отображения карты
type MapDataResponse struct {
	Points  []entity.Point `json:"points"`
	Routes  []entity.Route `json:"routes"`
	Center  entity.Point   `json:"center"`
	Zoom    int            `json:"zoom"`
	Source  string         `json:"source"`
	Updated string         `json:"updated"`
}

// MapUpdateRequest — обновление данных карты (например, после запроса к API)
type MapUpdateRequest struct {
	Source string `json:"source" validate:"required"` // "osm" или "serpapi"
	Limit  int    `json:"limit,omitempty"`
}
