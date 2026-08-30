package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lgomesroc/bus-track-go/domain"
)

var ErrLineNotFound = errors.New("line not found")

type LineRepository struct {
	db *sql.DB
}

func NewLineRepository(db *sql.DB) *LineRepository {
	return &LineRepository{
		db: db,
	}
}

func (r *LineRepository) FindAll() ([]domain.Line, error) {
	rows, err := r.db.Query(`
		SELECT
			ID,
			CODE,
			ORIGIN,
			DESTINATION
		FROM LINE
		ORDER BY ID
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to find lines: %w", err)
	}
	defer rows.Close()

	var lines []domain.Line

	for rows.Next() {
		var line domain.Line

		err := rows.Scan(
			&line.ID,
			&line.Code,
			&line.Origin,
			&line.Destination,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan line: %w", err)
		}

		lines = append(lines, line)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate lines: %w", err)
	}

	return lines, nil
}

func (r *LineRepository) FindByID(id int) (*domain.Line, error) {
	var line domain.Line

	err := r.db.QueryRow(`
		SELECT
			ID,
			CODE,
			ORIGIN,
			DESTINATION
		FROM LINE
		WHERE ID = :1
	`, id).Scan(
		&line.ID,
		&line.Code,
		&line.Origin,
		&line.Destination,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find line by id: %w", err)
	}

	return &line, nil
}

func (r *LineRepository) Create(line domain.Line) (*domain.Line, error) {
	result, err := r.db.Exec(`
		INSERT INTO LINE (
			CODE,
			ORIGIN,
			DESTINATION
		)
		VALUES (
			:1,
			:2,
			:3
		)
	`,
		line.Code,
		line.Origin,
		line.Destination,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create line: %w", err)
	}

	id, err := result.LastInsertId()
	if err == nil && id > 0 {
		return r.FindByID(int(id))
	}

	createdLine, err := r.findLastInsertedLine(line)
	if err != nil {
		return nil, err
	}

	return createdLine, nil
}

func (r *LineRepository) findLastInsertedLine(line domain.Line) (*domain.Line, error) {
	var createdLine domain.Line

	err := r.db.QueryRow(`
		SELECT
			ID,
			CODE,
			ORIGIN,
			DESTINATION
		FROM LINE
		WHERE CODE = :1
		  AND ORIGIN = :2
		  AND DESTINATION = :3
		ORDER BY ID DESC
		FETCH FIRST 1 ROW ONLY
	`,
		line.Code,
		line.Origin,
		line.Destination,
	).Scan(
		&createdLine.ID,
		&createdLine.Code,
		&createdLine.Origin,
		&createdLine.Destination,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created line: %w", err)
	}

	return &createdLine, nil
}

func (r *LineRepository) Update(id int, line domain.Line) (*domain.Line, error) {
	result, err := r.db.Exec(`
		UPDATE LINE
		SET
			CODE = :1,
			ORIGIN = :2,
			DESTINATION = :3
		WHERE ID = :4
	`,
		line.Code,
		line.Origin,
		line.Destination,
		id,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update line: %w", err)
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

func (r *LineRepository) Delete(id int) error {
	result, err := r.db.Exec(`
		DELETE FROM LINE
		WHERE ID = :1
	`, id)

	if err != nil {
		return fmt.Errorf("failed to delete line: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return ErrLineNotFound
	}

	return nil
}
