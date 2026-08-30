package domain

type Trip struct {
	ID         int    `json:"id"`
	LineID     int    `json:"lineId"`
	BusID      int    `json:"busId"`
	TripDate   string `json:"tripDate"`
	TripTime   string `json:"tripTime"`
	Passengers int    `json:"passengers"`
}