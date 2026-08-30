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

type LineRepository interface {
	FindAll() ([]domain.Line, error)
	FindByID(id int) (*domain.Line, error)
	Create(line domain.Line) (*domain.Line, error)
	Update(id int, line domain.Line) (*domain.Line, error)
	Delete(id int) error
}

type TripRepository interface {
	FindAll() ([]domain.Trip, error)
	FindByID(id int) (*domain.Trip, error)
	Create(trip domain.Trip) (*domain.Trip, error)
	Update(id int, trip domain.Trip) (*domain.Trip, error)
	Delete(id int) error
	AveragePassengersByLine(lineID int) (float64, error)
}

func main() {
	db, err := database.NewOracleConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	busRepository := repository.NewBusRepository(db)
	lineRepository := repository.NewLineRepository(db)
	tripRepository := repository.NewTripRepository(db)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)

	mux.HandleFunc("/api/buses", busesHandler(busRepository))
	mux.HandleFunc("/api/buses/", busByIDHandler(busRepository))

	mux.HandleFunc("/api/lines", linesHandler(lineRepository))
	mux.HandleFunc("/api/lines/", lineByIDHandler(lineRepository, tripRepository))

	mux.HandleFunc("/api/trips", tripsHandler(tripRepository))
	mux.HandleFunc("/api/trips/", tripByIDHandler(tripRepository))

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

func linesHandler(lineRepository LineRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			lines, err := lineRepository.FindAll()
			if err != nil {
				http.Error(w, "failed to find lines", http.StatusInternalServerError)
				return
			}

			writeJSON(w, http.StatusOK, lines)

		case http.MethodPost:
			var line domain.Line

			if err := decodeJSON(r, &line); err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}

			if err := validateLine(line); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			createdLine, err := lineRepository.Create(line)
			if err != nil {
				http.Error(w, "failed to create line", http.StatusInternalServerError)
				return
			}

			writeJSON(w, http.StatusCreated, createdLine)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func lineByIDHandler(
	lineRepository LineRepository,
	tripRepository TripRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/lines/")

		if strings.HasSuffix(path, "/average-passengers") {
			idString := strings.TrimSuffix(path, "/average-passengers")

			id, err := strconv.Atoi(idString)
			if err != nil || id <= 0 {
				http.Error(w, "invalid line id", http.StatusBadRequest)
				return
			}

			if r.Method != http.MethodGet {
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}

			average, err := tripRepository.AveragePassengersByLine(id)
			if err != nil {
				http.Error(w, "failed to calculate average passengers", http.StatusInternalServerError)
				return
			}

			writeJSON(w, http.StatusOK, map[string]interface{}{
				"lineId":            id,
				"averagePassengers": average,
			})

			return
		}

		id, err := strconv.Atoi(path)
		if err != nil || id <= 0 {
			http.Error(w, "invalid line id", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			line, err := lineRepository.FindByID(id)
			if err != nil {
				http.Error(w, "failed to find line", http.StatusInternalServerError)
				return
			}

			if line == nil {
				http.Error(w, "line not found", http.StatusNotFound)
				return
			}

			writeJSON(w, http.StatusOK, line)

		case http.MethodPut:
			var line domain.Line

			if err := decodeJSON(r, &line); err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}

			if err := validateLine(line); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			updatedLine, err := lineRepository.Update(id, line)
			if err != nil {
				http.Error(w, "failed to update line", http.StatusInternalServerError)
				return
			}

			if updatedLine == nil {
				http.Error(w, "line not found", http.StatusNotFound)
				return
			}

			writeJSON(w, http.StatusOK, updatedLine)

		case http.MethodDelete:
			err := lineRepository.Delete(id)
			if err != nil {
				if errors.Is(err, repository.ErrLineNotFound) {
					http.Error(w, "line not found", http.StatusNotFound)
					return
				}

				http.Error(w, "failed to delete line", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func tripsHandler(tripRepository TripRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			trips, err := tripRepository.FindAll()
			if err != nil {
				http.Error(w, "failed to find trips", http.StatusInternalServerError)
				return
			}

			writeJSON(w, http.StatusOK, trips)

		case http.MethodPost:
			var trip domain.Trip

			if err := decodeJSON(r, &trip); err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}

			if err := validateTrip(trip); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			createdTrip, err := tripRepository.Create(trip)
			if err != nil {
				log.Printf("failed to create trip: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			writeJSON(w, http.StatusCreated, createdTrip)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func tripByIDHandler(tripRepository TripRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseID(r.URL.Path, "/api/trips/")
		if err != nil {
			http.Error(w, "invalid trip id", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			trip, err := tripRepository.FindByID(id)
			if err != nil {
				http.Error(w, "failed to find trip", http.StatusInternalServerError)
				return
			}

			if trip == nil {
				http.Error(w, "trip not found", http.StatusNotFound)
				return
			}

			writeJSON(w, http.StatusOK, trip)

		case http.MethodPut:
			var trip domain.Trip

			if err := decodeJSON(r, &trip); err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}

			if err := validateTrip(trip); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			updatedTrip, err := tripRepository.Update(id, trip)
			if err != nil {
				http.Error(w, "failed to update trip", http.StatusInternalServerError)
				return
			}

			if updatedTrip == nil {
				http.Error(w, "trip not found", http.StatusNotFound)
				return
			}

			writeJSON(w, http.StatusOK, updatedTrip)

		case http.MethodDelete:
			err := tripRepository.Delete(id)
			if err != nil {
				if errors.Is(err, repository.ErrTripNotFound) {
					http.Error(w, "trip not found", http.StatusNotFound)
					return
				}

				http.Error(w, "failed to delete trip", http.StatusInternalServerError)
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

func validateLine(line domain.Line) error {
	if strings.TrimSpace(line.Code) == "" {
		return errors.New("code is required")
	}

	if strings.TrimSpace(line.Origin) == "" {
		return errors.New("origin is required")
	}

	if strings.TrimSpace(line.Destination) == "" {
		return errors.New("destination is required")
	}

	return nil
}

func validateTrip(trip domain.Trip) error {
	if trip.LineID <= 0 {
		return errors.New("lineId must be greater than zero")
	}

	if trip.BusID <= 0 {
		return errors.New("busId must be greater than zero")
	}

	if strings.TrimSpace(trip.TripDate) == "" {
		return errors.New("tripDate is required")
	}

	if strings.TrimSpace(trip.TripTime) == "" {
		return errors.New("tripTime is required")
	}

	if trip.Passengers < 0 {
		return errors.New("passengers must be greater than or equal to zero")
	}

	return nil
}

func parseBusID(path string) (int, error) {
	return parseID(path, "/api/buses/")
}

func parseID(path string, prefix string) (int, error) {
	idString := strings.TrimPrefix(path, prefix)

	if idString == "" || strings.Contains(idString, "/") {
		return 0, errors.New("invalid id")
	}

	id, err := strconv.Atoi(idString)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid id")
	}

	return id, nil
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
