package mapsapi

type GeoResponse struct {
	DisplayName string  `json:"display_name"`
	Lat         float64 `json:"lat,string"`
	Lon         float64 `json:"lon,string"`
	Type        string  `json:"type"`
}
