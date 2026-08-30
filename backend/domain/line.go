package domain

type Line struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}