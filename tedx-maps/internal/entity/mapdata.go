package entity

// MapData агрегирует данные для визуализации карты
type MapData struct {
	Points  []Point `json:"points"`
	Routes  []Route `json:"routes"`
	Center  Point   `json:"center"`  // центральная точка отображения
	Zoom    int     `json:"zoom"`    // уровень масштаба
	Source  string  `json:"source"`  // "osm" или "serpapi"
	Updated string  `json:"updated"` // время последнего обновления
}
