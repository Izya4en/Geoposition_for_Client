package entity

type Terminal struct {
	ID            string  `json:"id"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	AvailableCash float64 `json:"available_cash"`
}
