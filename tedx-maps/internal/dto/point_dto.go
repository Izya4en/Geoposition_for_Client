package dto

// PointCreateRequest — структура для создания новой точки
type PointCreateRequest struct {
	Name      string  `json:"name" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Type      string  `json:"type" validate:"required"`
}

// PointUpdateRequest — обновление существующей точки
type PointUpdateRequest struct {
	Name      *string  `json:"name,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Type      *string  `json:"type,omitempty"`
}

// PointResponse — стандартный ответ API при получении точки
type PointResponse struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Type      string  `json:"type"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
