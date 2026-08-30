package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lgomesroc/bus-track-go/domain"
)

var ErrTripNotFound = errors.New("trip not found")

type TripRepository struct {
	db *sql.DB
}

func NewTripRepository(db *sql.DB) *TripRepository {
	return &TripRepository{
		db: db,
	}
}

func (r *TripRepository) FindAll() ([]domain.Trip, error) {
	rows, err := r.db.Query(`
		SELECT
			ID,
			LINE_ID,
			BUS_ID,
			TO_CHAR(TRIP_DATE, 'YYYY-MM-DD'),
			TRIP_TIME,
			PASSENGERS
		FROM TRIP
		ORDER BY TRIP_DATE, TRIP_TIME, ID
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to find trips: %w", err)
	}
	defer rows.Close()

	var trips []domain.Trip

	for rows.Next() {
		var trip domain.Trip

		err := rows.Scan(
			&trip.ID,
			&trip.LineID,
			&trip.BusID,
			&trip.TripDate,
			&trip.TripTime,
			&trip.Passengers,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan trip: %w", err)
		}

		trips = append(trips, trip)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate trips: %w", err)
	}

	return trips, nil
}

func (r *TripRepository) FindByID(id int) (*domain.Trip, error) {
	var trip domain.Trip

	err := r.db.QueryRow(`
		SELECT
			ID,
			LINE_ID,
			BUS_ID,
			TO_CHAR(TRIP_DATE, 'YYYY-MM-DD'),
			TRIP_TIME,
			PASSENGERS
		FROM TRIP
		WHERE ID = :1
	`, id).Scan(
		&trip.ID,
		&trip.LineID,
		&trip.BusID,
		&trip.TripDate,
		&trip.TripTime,
		&trip.Passengers,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find trip by id: %w", err)
	}

	return &trip, nil
}

func (r *TripRepository) Create(trip domain.Trip) (*domain.Trip, error) {
	tripDate, err := time.Parse("2006-01-02", trip.TripDate)
	if err != nil {
		return nil, fmt.Errorf("invalid trip date: %w", err)
	}

	_, err = r.db.Exec(`
		INSERT INTO TRIP (
			LINE_ID,
			BUS_ID,
			TRIP_DATE,
			TRIP_TIME,
			PASSENGERS
		)
		VALUES (
			:1,
			:2,
			:3,
			:4,
			:5
		)
	`,
		trip.LineID,
		trip.BusID,
		tripDate,
		trip.TripTime,
		trip.Passengers,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create trip: %w", err)
	}

	var createdTrip domain.Trip

	err = r.db.QueryRow(`
		SELECT
			ID,
			LINE_ID,
			BUS_ID,
			TO_CHAR(TRIP_DATE, 'YYYY-MM-DD'),
			TRIP_TIME,
			PASSENGERS
		FROM TRIP
		WHERE LINE_ID = :1
		  AND BUS_ID = :2
		  AND TRUNC(TRIP_DATE) = TRUNC(:3)
		  AND TRIP_TIME = :4
		  AND PASSENGERS = :5
		ORDER BY ID DESC
		FETCH FIRST 1 ROW ONLY
	`,
		trip.LineID,
		trip.BusID,
		tripDate,
		trip.TripTime,
		trip.Passengers,
	).Scan(
		&createdTrip.ID,
		&createdTrip.LineID,
		&createdTrip.BusID,
		&createdTrip.TripDate,
		&createdTrip.TripTime,
		&createdTrip.Passengers,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created trip: %w", err)
	}

	return &createdTrip, nil
}

func (r *TripRepository) Update(id int, trip domain.Trip) (*domain.Trip, error) {
	tripDate, err := time.Parse("2006-01-02", trip.TripDate)
	if err != nil {
		return nil, fmt.Errorf("invalid trip date: %w", err)
	}

	result, err := r.db.Exec(`
		UPDATE TRIP
		SET
			LINE_ID = :1,
			BUS_ID = :2,
			TRIP_DATE = :3,
			TRIP_TIME = :4,
			PASSENGERS = :5
		WHERE ID = :6
	`,
		trip.LineID,
		trip.BusID,
		tripDate,
		trip.TripTime,
		trip.Passengers,
		id,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update trip: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return nil, nil
	}

	return r.FindByID(id)
}

func (r *TripRepository) Delete(id int) error {
	result, err := r.db.Exec(`
		DELETE FROM TRIP
		WHERE ID = :1
	`, id)

	if err != nil {
		return fmt.Errorf("failed to delete trip: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return ErrTripNotFound
	}

	return nil
}

func (r *TripRepository) AveragePassengersByLine(lineID int) (float64, error) {
	var average sql.NullFloat64

	err := r.db.QueryRow(`
		SELECT AVG(PASSENGERS)
		FROM TRIP
		WHERE LINE_ID = :1
	`, lineID).Scan(&average)

	if err != nil {
		return 0, fmt.Errorf("failed to calculate average passengers: %w", err)
	}

	if !average.Valid {
		return 0, nil
	}

	return average.Float64, nil
}
