package dto

// RouteCreateRequest — создание маршрута между двумя точками
type RouteCreateRequest struct {
	FromPointID int64 `json:"from_point_id" validate:"required"`
	ToPointID   int64 `json:"to_point_id" validate:"required"`
}

// RouteResponse — ответ при получении маршрута
type RouteResponse struct {
	ID         int64   `json:"id"`
	FromPoint  int64   `json:"from_point"`
	ToPoint    int64   `json:"to_point"`
	DistanceKM float64 `json:"distance_km"`
	Duration   string  `json:"duration"`
	Path       string  `json:"path"`
}
