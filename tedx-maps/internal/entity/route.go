package entity

// Route описывает маршрут между двумя точками
type Route struct {
	ID         int64   `db:"id" json:"id"`
	FromPoint  int64   `db:"from_point" json:"from_point"`
	ToPoint    int64   `db:"to_point" json:"to_point"`
	DistanceKM float64 `db:"distance_km" json:"distance_km"`
	Duration   string  `db:"duration" json:"duration"` // пример: "15m 30s"
	Path       string  `db:"path" json:"path"`         // JSON-координаты маршрута
	CreatedAt  string  `db:"created_at" json:"created_at"`
}
