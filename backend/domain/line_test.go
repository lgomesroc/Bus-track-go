package domain

import (
	"encoding/json"
	"testing"
)

func TestLineJSON(t *testing.T) {
	line := Line{
		ID:          1,
		Code:        "001",
		Origin:      "Nova Iguaçu",
		Destination: "Centro",
	}

	data, err := json.Marshal(line)
	if err != nil {
		t.Fatalf("erro ao converter Line para JSON: %v", err)
	}

	expected := `{"id":1,"code":"001","origin":"Nova Iguaçu","destination":"Centro"}`

	if string(data) != expected {
		t.Fatalf("JSON inesperado: %s", data)
	}
}

func TestLineFromJSON(t *testing.T) {
	data := []byte(`{
		"id": 2,
		"code": "002",
		"origin": "Belford Roxo",
		"destination": "Centro"
	}`)

	var line Line

	err := json.Unmarshal(data, &line)
	if err != nil {
		t.Fatalf("erro ao converter JSON para Line: %v", err)
	}

	if line.ID != 2 {
		t.Fatalf("ID inesperado: %d", line.ID)
	}

	if line.Code != "002" {
		t.Fatalf("Code inesperado: %s", line.Code)
	}

	if line.Origin != "Belford Roxo" {
		t.Fatalf("Origin inesperado: %s", line.Origin)
	}

	if line.Destination != "Centro" {
		t.Fatalf("Destination inesperado: %s", line.Destination)
	}
}
