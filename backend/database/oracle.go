package database

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func NewOracleConnection() (*sql.DB, error) {
	dsn := `user="system" password="BusTrack123" connectString="localhost:1521/FREEPDB1"`

	db, err := sql.Open("godror", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open Oracle connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping Oracle database: %w", err)
	}

	return db, nil
}