package domain

type Bus struct {
	ID           int    `json:"id"`
	Prefix       string `json:"prefix"`
	LicensePlate string `json:"licensePlate"`
	Model        string `json:"model"`
	Capacity     int    `json:"capacity"`
	Status       string `json:"status"`
}