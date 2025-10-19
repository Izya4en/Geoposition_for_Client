package entity

// MapData агрегирует данные для визуализации карты
type MapData struct {
	Points  []Point                `json:"points"`
	Routes  []Route                `json:"routes"`
	Center  Point                  `json:"center"`
	Zoom    int                    `json:"zoom"`
	Source  string                 `json:"source"`
	Updated string                 `json:"updated"`
	Data    map[string]interface{} `json:"data,omitempty"` // ← добавить вот это
}
