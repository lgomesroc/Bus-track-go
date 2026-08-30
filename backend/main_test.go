package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lgomesroc/bus-track-go/domain"
	"github.com/lgomesroc/bus-track-go/repository"
)

type fakeBusRepository struct {
	buses           []domain.Bus
	findAllErr      error
	findByIDErr     error
	createErr       error
	updateErr       error
	deleteErr       error
	busToFind       *domain.Bus
	busToCreate     *domain.Bus
	busToUpdate     *domain.Bus
	deleteErrToSend error
}

func (f *fakeBusRepository) FindAll() ([]domain.Bus, error) {
	if f.findAllErr != nil {
		return nil, f.findAllErr
	}

	return f.buses, nil
}

func (f *fakeBusRepository) FindByID(id int) (*domain.Bus, error) {
	if f.findByIDErr != nil {
		return nil, f.findByIDErr
	}

	return f.busToFind, nil
}

func (f *fakeBusRepository) Create(bus domain.Bus) (*domain.Bus, error) {
	if f.createErr != nil {
		return nil, f.createErr
	}

	if f.busToCreate != nil {
		return f.busToCreate, nil
	}

	bus.ID = 1
	return &bus, nil
}

func (f *fakeBusRepository) Update(id int, bus domain.Bus) (*domain.Bus, error) {
	if f.updateErr != nil {
		return nil, f.updateErr
	}

	return f.busToUpdate, nil
}

func (f *fakeBusRepository) Delete(id int) error {
	if f.deleteErrToSend != nil {
		return f.deleteErrToSend
	}

	return nil
}

type fakeLineRepository struct {
	lines           []domain.Line
	findAllErr      error
	findByIDErr     error
	createErr       error
	updateErr       error
	deleteErr       error
	lineToFind      *domain.Line
	lineToCreate    *domain.Line
	lineToUpdate    *domain.Line
	deleteErrToSend error
}

func (f *fakeLineRepository) FindAll() ([]domain.Line, error) {
	if f.findAllErr != nil {
		return nil, f.findAllErr
	}

	return f.lines, nil
}

func (f *fakeLineRepository) FindByID(id int) (*domain.Line, error) {
	if f.findByIDErr != nil {
		return nil, f.findByIDErr
	}

	return f.lineToFind, nil
}

func (f *fakeLineRepository) Create(line domain.Line) (*domain.Line, error) {
	if f.createErr != nil {
		return nil, f.createErr
	}

	if f.lineToCreate != nil {
		return f.lineToCreate, nil
	}

	line.ID = 1
	return &line, nil
}

func (f *fakeLineRepository) Update(id int, line domain.Line) (*domain.Line, error) {
	if f.updateErr != nil {
		return nil, f.updateErr
	}

	return f.lineToUpdate, nil
}

func (f *fakeLineRepository) Delete(id int) error {
	if f.deleteErrToSend != nil {
		return f.deleteErrToSend
	}

	return nil
}

type fakeTripRepository struct {
	trips                   []domain.Trip
	findAllErr              error
	findByIDErr             error
	createErr               error
	updateErr               error
	deleteErr               error
	averagePassengersErr    error
	averagePassengersResult float64
	tripToFind              *domain.Trip
	tripToCreate            *domain.Trip
	tripToUpdate            *domain.Trip
	deleteErrToSend         error
}

func (f *fakeTripRepository) FindAll() ([]domain.Trip, error) {
	if f.findAllErr != nil {
		return nil, f.findAllErr
	}

	return f.trips, nil
}

func (f *fakeTripRepository) FindByID(id int) (*domain.Trip, error) {
	if f.findByIDErr != nil {
		return nil, f.findByIDErr
	}

	return f.tripToFind, nil
}

func (f *fakeTripRepository) Create(trip domain.Trip) (*domain.Trip, error) {
	if f.createErr != nil {
		return nil, f.createErr
	}

	if f.tripToCreate != nil {
		return f.tripToCreate, nil
	}

	trip.ID = 1
	return &trip, nil
}

func (f *fakeTripRepository) Update(id int, trip domain.Trip) (*domain.Trip, error) {
	if f.updateErr != nil {
		return nil, f.updateErr
	}

	return f.tripToUpdate, nil
}

func (f *fakeTripRepository) Delete(id int) error {
	if f.deleteErrToSend != nil {
		return f.deleteErrToSend
	}

	return nil
}

func (f *fakeTripRepository) AveragePassengersByLine(lineID int) (float64, error) {
	if f.averagePassengersErr != nil {
		return 0, f.averagePassengersErr
	}

	return f.averagePassengersResult, nil
}

func validBus() domain.Bus {
	return domain.Bus{
		ID:           21,
		Prefix:       "001",
		LicensePlate: "ABC1D23",
		Model:        "Mercedes-Benz",
		Capacity:     45,
		Status:       "active",
	}
}

func validLine() domain.Line {
	return domain.Line{
		ID:          1,
		Code:        "001",
		Origin:      "Nova Iguaçu",
		Destination: "Centro",
	}
}

func validTrip() domain.Trip {
	return domain.Trip{
		ID:         1,
		LineID:     1,
		BusID:      21,
		TripDate:   "2026-08-29",
		TripTime:   "08:00",
		Passengers: 38,
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), `"status":"ok"`) {
		t.Errorf("expected response body to contain status ok, got %s", rec.Body.String())
	}
}

func TestHealthHandlerMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d, got %d", http.StatusMethodNotAllowed, rec.Code)
	}
}

func TestValidateBus(t *testing.T) {
	tests := []struct {
		name    string
		bus     domain.Bus
		wantErr bool
	}{
		{
			name:    "valid bus",
			bus:     validBus(),
			wantErr: false,
		},
		{
			name: "missing prefix",
			bus: domain.Bus{
				LicensePlate: "ABC-1234",
				Model:        "Mercedes-Benz",
				Capacity:     40,
				Status:       "ACTIVE",
			},
			wantErr: true,
		},
		{
			name: "missing license plate",
			bus: domain.Bus{
				Prefix:   "BUS-001",
				Model:    "Mercedes-Benz",
				Capacity: 40,
				Status:   "ACTIVE",
			},
			wantErr: true,
		},
		{
			name: "missing model",
			bus: domain.Bus{
				Prefix:       "BUS-001",
				LicensePlate: "ABC-1234",
				Capacity:     40,
				Status:       "ACTIVE",
			},
			wantErr: true,
		},
		{
			name: "invalid capacity",
			bus: domain.Bus{
				Prefix:       "BUS-001",
				LicensePlate: "ABC-1234",
				Model:        "Mercedes-Benz",
				Capacity:     0,
				Status:       "ACTIVE",
			},
			wantErr: true,
		},
		{
			name: "missing status",
			bus: domain.Bus{
				Prefix:       "BUS-001",
				LicensePlate: "ABC-1234",
				Model:        "Mercedes-Benz",
				Capacity:     40,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateBus(tt.bus)

			if (err != nil) != tt.wantErr {
				t.Errorf("validateBus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateLine(t *testing.T) {
	tests := []struct {
		name    string
		line    domain.Line
		wantErr bool
	}{
		{
			name:    "valid line",
			line:    validLine(),
			wantErr: false,
		},
		{
			name: "missing code",
			line: domain.Line{
				Origin:      "Nova Iguaçu",
				Destination: "Centro",
			},
			wantErr: true,
		},
		{
			name: "missing origin",
			line: domain.Line{
				Code:        "001",
				Destination: "Centro",
			},
			wantErr: true,
		},
		{
			name: "missing destination",
			line: domain.Line{
				Code:   "001",
				Origin: "Nova Iguaçu",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateLine(tt.line)

			if (err != nil) != tt.wantErr {
				t.Errorf("validateLine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateTrip(t *testing.T) {
	tests := []struct {
		name    string
		trip    domain.Trip
		wantErr bool
	}{
		{
			name:    "valid trip",
			trip:    validTrip(),
			wantErr: false,
		},
		{
			name: "invalid line id",
			trip: domain.Trip{
				LineID:     0,
				BusID:      21,
				TripDate:   "2026-08-29",
				TripTime:   "08:00",
				Passengers: 38,
			},
			wantErr: true,
		},
		{
			name: "invalid bus id",
			trip: domain.Trip{
				LineID:     1,
				BusID:      0,
				TripDate:   "2026-08-29",
				TripTime:   "08:00",
				Passengers: 38,
			},
			wantErr: true,
		},
		{
			name: "missing trip date",
			trip: domain.Trip{
				LineID:     1,
				BusID:      21,
				TripTime:   "08:00",
				Passengers: 38,
			},
			wantErr: true,
		},
		{
			name: "missing trip time",
			trip: domain.Trip{
				LineID:     1,
				BusID:      21,
				TripDate:   "2026-08-29",
				Passengers: 38,
			},
			wantErr: true,
		},
		{
			name: "negative passengers",
			trip: domain.Trip{
				LineID:     1,
				BusID:      21,
				TripDate:   "2026-08-29",
				TripTime:   "08:00",
				Passengers: -1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTrip(tt.trip)

			if (err != nil) != tt.wantErr {
				t.Errorf("validateTrip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseBusID(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantID  int
		wantErr bool
	}{
		{
			name:   "valid id",
			path:   "/api/buses/1",
			wantID: 1,
		},
		{
			name:   "valid id 10",
			path:   "/api/buses/10",
			wantID: 10,
		},
		{
			name:    "missing id",
			path:    "/api/buses/",
			wantErr: true,
		},
		{
			name:    "non numeric id",
			path:    "/api/buses/abc",
			wantErr: true,
		},
		{
			name:    "path with additional segment",
			path:    "/api/buses/1/details",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, err := parseBusID(tt.path)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseBusID() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && gotID != tt.wantID {
				t.Errorf("parseBusID() = %d, want %d", gotID, tt.wantID)
			}
		})
	}
}

func TestBusesHandlerGet(t *testing.T) {
	fakeRepository := &fakeBusRepository{
		buses: []domain.Bus{
			validBus(),
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/api/buses", nil)
	rec := httptest.NewRecorder()

	handler := busesHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "001") {
		t.Errorf("expected response to contain 001, got %s", rec.Body.String())
	}
}

func TestBusesHandlerGetError(t *testing.T) {
	fakeRepository := &fakeBusRepository{
		findAllErr: errors.New("database error"),
	}

	req := httptest.NewRequest(http.MethodGet, "/api/buses", nil)
	rec := httptest.NewRecorder()

	handler := busesHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}

func TestBusesHandlerPost(t *testing.T) {
	bus := validBus()

	body, err := json.Marshal(bus)
	if err != nil {
		t.Fatalf("failed to marshal bus: %v", err)
	}

	fakeRepository := &fakeBusRepository{
		busToCreate: &bus,
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/buses",
		bytes.NewReader(body),
	)

	rec := httptest.NewRecorder()

	handler := busesHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "001") {
		t.Errorf("expected response to contain 001, got %s", rec.Body.String())
	}
}

func TestBusesHandlerPostInvalidJSON(t *testing.T) {
	fakeRepository := &fakeBusRepository{}

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/buses",
		strings.NewReader(`{"prefix":`),
	)

	rec := httptest.NewRecorder()

	handler := busesHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestBusesHandlerPostInvalidBus(t *testing.T) {
	bus := domain.Bus{
		Prefix: "",
	}

	body, err := json.Marshal(bus)
	if err != nil {
		t.Fatalf("failed to marshal bus: %v", err)
	}

	fakeRepository := &fakeBusRepository{}

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/buses",
		bytes.NewReader(body),
	)

	rec := httptest.NewRecorder()

	handler := busesHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestBusesHandlerMethodNotAllowed(t *testing.T) {
	fakeRepository := &fakeBusRepository{}

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/buses",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := busesHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d, got %d", http.StatusMethodNotAllowed, rec.Code)
	}
}

func TestBusByIDHandlerGet(t *testing.T) {
	bus := validBus()

	fakeRepository := &fakeBusRepository{
		busToFind: &bus,
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/buses/21",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "001") {
		t.Errorf("expected response to contain 001, got %s", rec.Body.String())
	}
}

func TestBusByIDHandlerGetNotFound(t *testing.T) {
	fakeRepository := &fakeBusRepository{
		busToFind: nil,
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/buses/999",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestBusByIDHandlerGetError(t *testing.T) {
	fakeRepository := &fakeBusRepository{
		findByIDErr: errors.New("database error"),
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/buses/1",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}

func TestBusByIDHandlerGetInvalidID(t *testing.T) {
	fakeRepository := &fakeBusRepository{}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/buses/abc",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestBusByIDHandlerPut(t *testing.T) {
	bus := validBus()

	body, err := json.Marshal(bus)
	if err != nil {
		t.Fatalf("failed to marshal bus: %v", err)
	}

	fakeRepository := &fakeBusRepository{
		busToUpdate: &bus,
	}

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/buses/21",
		bytes.NewReader(body),
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "001") {
		t.Errorf("expected response to contain 001, got %s", rec.Body.String())
	}
}

func TestBusByIDHandlerPutNotFound(t *testing.T) {
	bus := validBus()

	body, err := json.Marshal(bus)
	if err != nil {
		t.Fatalf("failed to marshal bus: %v", err)
	}

	fakeRepository := &fakeBusRepository{
		busToUpdate: nil,
	}

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/buses/999",
		bytes.NewReader(body),
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestBusByIDHandlerPutInvalidBus(t *testing.T) {
	bus := domain.Bus{
		Prefix: "",
	}

	body, err := json.Marshal(bus)
	if err != nil {
		t.Fatalf("failed to marshal bus: %v", err)
	}

	fakeRepository := &fakeBusRepository{}

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/buses/1",
		bytes.NewReader(body),
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestBusByIDHandlerDelete(t *testing.T) {
	fakeRepository := &fakeBusRepository{}

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/buses/1",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}
}

func TestBusByIDHandlerDeleteNotFound(t *testing.T) {
	fakeRepository := &fakeBusRepository{
		deleteErrToSend: repository.ErrBusNotFound,
	}

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/buses/999",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestBusByIDHandlerDeleteError(t *testing.T) {
	fakeRepository := &fakeBusRepository{
		deleteErrToSend: errors.New("database error"),
	}

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/buses/1",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}

func TestBusByIDHandlerMethodNotAllowed(t *testing.T) {
	fakeRepository := &fakeBusRepository{}

	req := httptest.NewRequest(
		http.MethodPatch,
		"/api/buses/1",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d, got %d", http.StatusMethodNotAllowed, rec.Code)
	}
}

func TestLinesHandlerGet(t *testing.T) {
	fakeRepository := &fakeLineRepository{
		lines: []domain.Line{
			validLine(),
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/api/lines", nil)
	rec := httptest.NewRecorder()

	handler := linesHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "Nova Igua") {
		t.Errorf("expected response to contain line origin, got %s", rec.Body.String())
	}
}

func TestLinesHandlerPost(t *testing.T) {
	line := validLine()

	body, err := json.Marshal(line)
	if err != nil {
		t.Fatalf("failed to marshal line: %v", err)
	}

	fakeRepository := &fakeLineRepository{
		lineToCreate: &line,
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/lines",
		bytes.NewReader(body),
	)

	rec := httptest.NewRecorder()

	handler := linesHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), `"code":"001"`) {
		t.Errorf("expected response to contain line code, got %s", rec.Body.String())
	}
}

func TestLineByIDHandlerGet(t *testing.T) {
	line := validLine()

	fakeRepository := &fakeLineRepository{
		lineToFind: &line,
	}

	fakeTripRepository := &fakeTripRepository{}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/lines/1",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := lineByIDHandler(fakeRepository, fakeTripRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), `"code":"001"`) {
		t.Errorf("expected response to contain line code, got %s", rec.Body.String())
	}
}

func TestLineByIDHandlerGetNotFound(t *testing.T) {
	fakeRepository := &fakeLineRepository{
		lineToFind: nil,
	}

	fakeTripRepository := &fakeTripRepository{}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/lines/999",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := lineByIDHandler(fakeRepository, fakeTripRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestLineByIDHandlerPut(t *testing.T) {
	line := validLine()

	body, err := json.Marshal(line)
	if err != nil {
		t.Fatalf("failed to marshal line: %v", err)
	}

	fakeRepository := &fakeLineRepository{
		lineToUpdate: &line,
	}

	fakeTripRepository := &fakeTripRepository{}

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/lines/1",
		bytes.NewReader(body),
	)

	rec := httptest.NewRecorder()

	handler := lineByIDHandler(fakeRepository, fakeTripRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), `"code":"001"`) {
		t.Errorf("expected response to contain line code, got %s", rec.Body.String())
	}
}

func TestLineByIDHandlerDelete(t *testing.T) {
	fakeRepository := &fakeLineRepository{}
	fakeTripRepository := &fakeTripRepository{}

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/lines/1",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := lineByIDHandler(fakeRepository, fakeTripRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}
}

func TestTripsHandlerGet(t *testing.T) {
	fakeRepository := &fakeTripRepository{
		trips: []domain.Trip{
			validTrip(),
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/api/trips", nil)
	rec := httptest.NewRecorder()

	handler := tripsHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), `"passengers":38`) {
		t.Errorf("expected response to contain passenger count, got %s", rec.Body.String())
	}
}

func TestTripsHandlerPost(t *testing.T) {
	trip := validTrip()

	body, err := json.Marshal(trip)
	if err != nil {
		t.Fatalf("failed to marshal trip: %v", err)
	}

	fakeRepository := &fakeTripRepository{
		tripToCreate: &trip,
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/trips",
		bytes.NewReader(body),
	)

	rec := httptest.NewRecorder()

	handler := tripsHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), `"passengers":38`) {
		t.Errorf("expected response to contain passenger count, got %s", rec.Body.String())
	}
}

func TestTripsHandlerPostInvalidTrip(t *testing.T) {
	trip := domain.Trip{
		LineID: 0,
		BusID:  21,
	}

	body, err := json.Marshal(trip)
	if err != nil {
		t.Fatalf("failed to marshal trip: %v", err)
	}

	fakeRepository := &fakeTripRepository{}

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/trips",
		bytes.NewReader(body),
	)

	rec := httptest.NewRecorder()

	handler := tripsHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestTripByIDHandlerGet(t *testing.T) {
	trip := validTrip()

	fakeRepository := &fakeTripRepository{
		tripToFind: &trip,
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/trips/1",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := tripByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), `"passengers":38`) {
		t.Errorf("expected response to contain passenger count, got %s", rec.Body.String())
	}
}

func TestTripByIDHandlerGetNotFound(t *testing.T) {
	fakeRepository := &fakeTripRepository{
		tripToFind: nil,
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/trips/999",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := tripByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestTripByIDHandlerPut(t *testing.T) {
	trip := validTrip()

	body, err := json.Marshal(trip)
	if err != nil {
		t.Fatalf("failed to marshal trip: %v", err)
	}

	fakeRepository := &fakeTripRepository{
		tripToUpdate: &trip,
	}

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/trips/1",
		bytes.NewReader(body),
	)

	rec := httptest.NewRecorder()

	handler := tripByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), `"passengers":38`) {
		t.Errorf("expected response to contain passenger count, got %s", rec.Body.String())
	}
}

func TestTripByIDHandlerDelete(t *testing.T) {
	fakeRepository := &fakeTripRepository{}

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/trips/1",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := tripByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}
}

func TestLineAveragePassengers(t *testing.T) {
	fakeLineRepository := &fakeLineRepository{}

	fakeTripRepository := &fakeTripRepository{
		averagePassengersResult: 37.5,
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/lines/1/average-passengers",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := lineByIDHandler(fakeLineRepository, fakeTripRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), `"lineId":1`) {
		t.Errorf("expected response to contain line id, got %s", rec.Body.String())
	}

	if !strings.Contains(rec.Body.String(), `"averagePassengers":37.5`) {
		t.Errorf("expected response to contain average passengers, got %s", rec.Body.String())
	}
}

func TestLineAveragePassengersError(t *testing.T) {
	fakeLineRepository := &fakeLineRepository{}

	fakeTripRepository := &fakeTripRepository{
		averagePassengersErr: errors.New("database error"),
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/lines/1/average-passengers",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := lineByIDHandler(fakeLineRepository, fakeTripRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusInternalServerError,
			rec.Code,
		)
	}
}
