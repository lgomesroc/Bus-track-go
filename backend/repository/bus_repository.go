package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lgomesroc/bus-track-go/domain"
)

var ErrBusNotFound = errors.New("bus not found")

type BusRepository struct {
	db *sql.DB
}

func NewBusRepository(db *sql.DB) *BusRepository {
	return &BusRepository{
		db: db,
	}
}

func (r *BusRepository) FindAll() ([]domain.Bus, error) {
	rows, err := r.db.Query(`
		SELECT
			ID,
			PREFIX,
			LICENSE_PLATE,
			MODEL,
			CAPACITY,
			STATUS
		FROM BUS
		ORDER BY ID
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to find buses: %w", err)
	}
	defer rows.Close()

	var buses []domain.Bus

	for rows.Next() {
		var bus domain.Bus

		err := rows.Scan(
			&bus.ID,
			&bus.Prefix,
			&bus.LicensePlate,
			&bus.Model,
			&bus.Capacity,
			&bus.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan bus: %w", err)
		}

		buses = append(buses, bus)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate buses: %w", err)
	}

	return buses, nil
}

func (r *BusRepository) FindByID(id int) (*domain.Bus, error) {
	var bus domain.Bus

	err := r.db.QueryRow(`
		SELECT
			ID,
			PREFIX,
			LICENSE_PLATE,
			MODEL,
			CAPACITY,
			STATUS
		FROM BUS
		WHERE ID = :1
	`, id).Scan(
		&bus.ID,
		&bus.Prefix,
		&bus.LicensePlate,
		&bus.Model,
		&bus.Capacity,
		&bus.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find bus by id: %w", err)
	}

	return &bus, nil
}

func (r *BusRepository) Create(bus domain.Bus) (*domain.Bus, error) {
	result, err := r.db.Exec(`
		INSERT INTO BUS (
			PREFIX,
			LICENSE_PLATE,
			MODEL,
			CAPACITY,
			STATUS
		)
		VALUES (
			:1,
			:2,
			:3,
			:4,
			:5
		)
	`,
		bus.Prefix,
		bus.LicensePlate,
		bus.Model,
		bus.Capacity,
		bus.Status,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create bus: %w", err)
	}

	id, err := result.LastInsertId()
	if err == nil && id > 0 {
		return r.FindByID(int(id))
	}

	createdBus, err := r.findLastInsertedBus(bus)
	if err != nil {
		return nil, err
	}

	return createdBus, nil
}

func (r *BusRepository) findLastInsertedBus(bus domain.Bus) (*domain.Bus, error) {
	var createdBus domain.Bus

	err := r.db.QueryRow(`
		SELECT
			ID,
			PREFIX,
			LICENSE_PLATE,
			MODEL,
			CAPACITY,
			STATUS
		FROM BUS
		WHERE PREFIX = :1
		  AND LICENSE_PLATE = :2
		  AND MODEL = :3
		  AND CAPACITY = :4
		  AND STATUS = :5
		ORDER BY ID DESC
		FETCH FIRST 1 ROW ONLY
	`,
		bus.Prefix,
		bus.LicensePlate,
		bus.Model,
		bus.Capacity,
		bus.Status,
	).Scan(
		&createdBus.ID,
		&createdBus.Prefix,
		&createdBus.LicensePlate,
		&createdBus.Model,
		&createdBus.Capacity,
		&createdBus.Status,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created bus: %w", err)
	}

	return &createdBus, nil
}

func (r *BusRepository) Update(id int, bus domain.Bus) (*domain.Bus, error) {
	result, err := r.db.Exec(`
		UPDATE BUS
		SET
			PREFIX = :1,
			LICENSE_PLATE = :2,
			MODEL = :3,
			CAPACITY = :4,
			STATUS = :5
		WHERE ID = :6
	`,
		bus.Prefix,
		bus.LicensePlate,
		bus.Model,
		bus.Capacity,
		bus.Status,
		id,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update bus: %w", err)
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

func (r *BusRepository) Delete(id int) error {
	result, err := r.db.Exec(`
		DELETE FROM BUS
		WHERE ID = :1
	`, id)

	if err != nil {
		return fmt.Errorf("failed to delete bus: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return ErrBusNotFound
	}

	return nil
}
