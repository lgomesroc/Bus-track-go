package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/lgomesroc/bus-track-go/database"
	"github.com/lgomesroc/bus-track-go/domain"
	"github.com/lgomesroc/bus-track-go/repository"
)

type BusRepository interface {
	FindAll() ([]domain.Bus, error)
	FindByID(id int) (*domain.Bus, error)
	Create(bus domain.Bus) (*domain.Bus, error)
	Update(id int, bus domain.Bus) (*domain.Bus, error)
	Delete(id int) error
}

func main() {
	db, err := database.NewOracleConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	busRepository := repository.NewBusRepository(db)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/api/buses", busesHandler(busRepository))
	mux.HandleFunc("/api/buses/", busByIDHandler(busRepository))

	log.Println("BusTrack API running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", corsMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func busesHandler(busRepository BusRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			buses, err := busRepository.FindAll()
			if err != nil {
				http.Error(w, "failed to find buses", http.StatusInternalServerError)
				return
			}

			writeJSON(w, http.StatusOK, buses)

		case http.MethodPost:
			var bus domain.Bus

			if err := decodeJSON(r, &bus); err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}

			if err := validateBus(bus); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			createdBus, err := busRepository.Create(bus)
			if err != nil {
				http.Error(w, "failed to create bus", http.StatusInternalServerError)
				return
			}

			writeJSON(w, http.StatusCreated, createdBus)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func busByIDHandler(busRepository BusRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseBusID(r.URL.Path)
		if err != nil {
			http.Error(w, "invalid bus id", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			bus, err := busRepository.FindByID(id)
			if err != nil {
				http.Error(w, "failed to find bus", http.StatusInternalServerError)
				return
			}

			if bus == nil {
				http.Error(w, "bus not found", http.StatusNotFound)
				return
			}

			writeJSON(w, http.StatusOK, bus)

		case http.MethodPut:
			var bus domain.Bus

			if err := decodeJSON(r, &bus); err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}

			if err := validateBus(bus); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			updatedBus, err := busRepository.Update(id, bus)
			if err != nil {
				http.Error(w, "failed to update bus", http.StatusInternalServerError)
				return
			}

			if updatedBus == nil {
				http.Error(w, "bus not found", http.StatusNotFound)
				return
			}

			writeJSON(w, http.StatusOK, updatedBus)

		case http.MethodDelete:
			err := busRepository.Delete(id)
			if err != nil {
				if errors.Is(err, repository.ErrBusNotFound) {
					http.Error(w, "bus not found", http.StatusNotFound)
					return
				}

				http.Error(w, "failed to delete bus", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func validateBus(bus domain.Bus) error {
	if strings.TrimSpace(bus.Prefix) == "" {
		return errors.New("prefix is required")
	}

	if strings.TrimSpace(bus.LicensePlate) == "" {
		return errors.New("licensePlate is required")
	}

	if strings.TrimSpace(bus.Model) == "" {
		return errors.New("model is required")
	}

	if bus.Capacity <= 0 {
		return errors.New("capacity must be greater than zero")
	}

	if strings.TrimSpace(bus.Status) == "" {
		return errors.New("status is required")
	}

	return nil
}

func parseBusID(path string) (int, error) {
	idString := strings.TrimPrefix(path, "/api/buses/")

	if idString == "" || strings.Contains(idString, "/") {
		return 0, errors.New("invalid bus id")
	}

	return strconv.Atoi(idString)
}

func decodeJSON(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(target)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to write JSON response: %v", err)
	}
}
