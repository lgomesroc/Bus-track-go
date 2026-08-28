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

func validBus() domain.Bus {
	return domain.Bus{
		ID:           1,
		Prefix:       "BUS-001",
		LicensePlate: "ABC-1234",
		Model:        "Mercedes-Benz",
		Capacity:     40,
		Status:       "ACTIVE",
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
			name: "valid bus",
			bus: domain.Bus{
				Prefix:       "BUS-001",
				LicensePlate: "ABC-1234",
				Model:        "Mercedes-Benz",
				Capacity:     40,
				Status:       "ACTIVE",
			},
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

	if !strings.Contains(rec.Body.String(), "BUS-001") {
		t.Errorf("expected response to contain BUS-001, got %s", rec.Body.String())
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
		t.Errorf(
			"expected status %d, got %d",
			http.StatusInternalServerError,
			rec.Code,
		)
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
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler := busesHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "BUS-001") {
		t.Errorf("expected response to contain BUS-001, got %s", rec.Body.String())
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
		t.Errorf(
			"expected status %d, got %d",
			http.StatusMethodNotAllowed,
			rec.Code,
		)
	}
}

func TestBusByIDHandlerGet(t *testing.T) {
	bus := validBus()

	fakeRepository := &fakeBusRepository{
		busToFind: &bus,
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/buses/1",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "BUS-001") {
		t.Errorf("expected response to contain BUS-001, got %s", rec.Body.String())
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
		t.Errorf(
			"expected status %d, got %d",
			http.StatusInternalServerError,
			rec.Code,
		)
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
		"/api/buses/1",
		bytes.NewReader(body),
	)

	rec := httptest.NewRecorder()

	handler := busByIDHandler(fakeRepository)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "BUS-001") {
		t.Errorf("expected response to contain BUS-001, got %s", rec.Body.String())
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
		t.Errorf(
			"expected status %d, got %d",
			http.StatusInternalServerError,
			rec.Code,
		)
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
		t.Errorf(
			"expected status %d, got %d",
			http.StatusMethodNotAllowed,
			rec.Code,
		)
	}
}
