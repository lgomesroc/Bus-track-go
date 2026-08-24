package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/lgomesroc/bus-track-go/domain"
)

var buses = []domain.Bus{
	{
		ID:           1,
		Prefix:       "001",
		LicensePlate: "ABC1D23",
		Model:        "Mercedes-Benz",
		Capacity:     50,
		Status:       "active",
	},
	{
		ID:           2,
		Prefix:       "002",
		LicensePlate: "DEF4G56",
		Model:        "Volkswagen",
		Capacity:     45,
		Status:       "active",
	},
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/buses", busesHandler)
	http.HandleFunc("/api/buses/", busByIDHandler)

	http.ListenAndServe(":8080", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status": "ok",
	}

	json.NewEncoder(w).Encode(response)
}

func busesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getBuses(w)
	case http.MethodPost:
		createBus(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func getBuses(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(buses)
}

func createBus(w http.ResponseWriter, r *http.Request) {
	var bus domain.Bus

	err := json.NewDecoder(r.Body).Decode(&bus)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	bus.ID = nextBusID()

	buses = append(buses, bus)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(bus)
}

func nextBusID() int {
	maxID := 0

	for _, bus := range buses {
		if bus.ID > maxID {
			maxID = bus.ID
		}
	}

	return maxID + 1
}

func busByIDHandler(w http.ResponseWriter, r *http.Request) {
	idText := strings.TrimPrefix(r.URL.Path, "/api/buses/")

	id, err := strconv.Atoi(idText)
	if err != nil {
		http.Error(w, "invalid bus id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getBusByID(w, id)
	case http.MethodPut:
		updateBus(w, r, id)
	case http.MethodDelete:
		deleteBus(w, id)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func getBusByID(w http.ResponseWriter, id int) {
	for _, bus := range buses {
		if bus.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(bus)
			return
		}
	}

	http.Error(w, "bus not found", http.StatusNotFound)
}

func updateBus(w http.ResponseWriter, r *http.Request, id int) {
	var updatedBus domain.Bus

	err := json.NewDecoder(r.Body).Decode(&updatedBus)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	for index, bus := range buses {
		if bus.ID == id {
			updatedBus.ID = id
			buses[index] = updatedBus

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedBus)
			return
		}
	}

	http.Error(w, "bus not found", http.StatusNotFound)
}

func deleteBus(w http.ResponseWriter, id int) {
	for index, bus := range buses {
		if bus.ID == id {
			buses = append(buses[:index], buses[index+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "bus not found", http.StatusNotFound)
}
