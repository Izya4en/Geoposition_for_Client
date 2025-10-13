package entity

// Point описывает географическую точку на карте (терминал, место TEDx и т.д.)
type Point struct {
	ID        int64   `db:"id" json:"id"`
	Name      string  `db:"name" json:"name"`
	Latitude  float64 `db:"latitude" json:"latitude"`
	Longitude float64 `db:"longitude" json:"longitude"`
	Type      string  `db:"type" json:"type"` // "terminal", "event", "custom"
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedAt string  `db:"updated_at" json:"updated_at"`
}
