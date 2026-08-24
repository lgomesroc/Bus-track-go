package domain

import (
	"encoding/json"
	"testing"
)

func TestBusJSON(t *testing.T) {
	bus := Bus{
		ID:           1,
		Prefix:       "001",
		LicensePlate: "ABC1D23",
		Model:        "Mercedes-Benz",
		Capacity:     50,
		Status:       "active",
	}

	data, err := json.Marshal(bus)
	if err != nil {
		t.Fatalf("erro ao converter Bus para JSON: %v", err)
	}

	expected := `{"id":1,"prefix":"001","licensePlate":"ABC1D23","model":"Mercedes-Benz","capacity":50,"status":"active"}`

	if string(data) != expected {
		t.Fatalf("JSON inesperado: %s", data)
	}
}

func TestBusFromJSON(t *testing.T) {
	data := []byte(`{
		"id": 2,
		"prefix": "002",
		"licensePlate": "DEF4G56",
		"model": "Volkswagen",
		"capacity": 45,
		"status": "maintenance"
	}`)

	var bus Bus

	err := json.Unmarshal(data, &bus)
	if err != nil {
		t.Fatalf("erro ao converter JSON para Bus: %v", err)
	}

	if bus.ID != 2 {
		t.Fatalf("ID inesperado: %d", bus.ID)
	}

	if bus.Prefix != "002" {
		t.Fatalf("Prefix inesperado: %s", bus.Prefix)
	}

	if bus.LicensePlate != "DEF4G56" {
		t.Fatalf("LicensePlate inesperado: %s", bus.LicensePlate)
	}

	if bus.Model != "Volkswagen" {
		t.Fatalf("Model inesperado: %s", bus.Model)
	}

	if bus.Capacity != 45 {
		t.Fatalf("Capacity inesperado: %d", bus.Capacity)
	}

	if bus.Status != "maintenance" {
		t.Fatalf("Status inesperado: %s", bus.Status)
	}
}