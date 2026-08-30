package domain

import (
	"encoding/json"
	"testing"
)

func TestTripJSON(t *testing.T) {
	trip := Trip{
		ID:         1,
		LineID:     1,
		BusID:      21,
		TripDate:   "2026-08-29",
		TripTime:   "08:00",
		Passengers: 38,
	}

	data, err := json.Marshal(trip)
	if err != nil {
		t.Fatalf("erro ao converter Trip para JSON: %v", err)
	}

	expected := `{"id":1,"lineId":1,"busId":21,"tripDate":"2026-08-29","tripTime":"08:00","passengers":38}`

	if string(data) != expected {
		t.Fatalf("JSON inesperado: %s", data)
	}
}

func TestTripFromJSON(t *testing.T) {
	data := []byte(`{
		"id": 2,
		"lineId": 1,
		"busId": 21,
		"tripDate": "2026-08-29",
		"tripTime": "08:00",
		"passengers": 42
	}`)

	var trip Trip

	err := json.Unmarshal(data, &trip)
	if err != nil {
		t.Fatalf("erro ao converter JSON para Trip: %v", err)
	}

	if trip.ID != 2 {
		t.Fatalf("ID inesperado: %d", trip.ID)
	}

	if trip.LineID != 1 {
		t.Fatalf("LineID inesperado: %d", trip.LineID)
	}

	if trip.BusID != 21 {
		t.Fatalf("BusID inesperado: %d", trip.BusID)
	}

	if trip.TripDate != "2026-08-29" {
		t.Fatalf("TripDate inesperado: %s", trip.TripDate)
	}

	if trip.TripTime != "08:00" {
		t.Fatalf("TripTime inesperado: %s", trip.TripTime)
	}

	if trip.Passengers != 42 {
		t.Fatalf("Passengers inesperado: %d", trip.Passengers)
	}
}
