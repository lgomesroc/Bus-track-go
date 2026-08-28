package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/godror/godror"
)

func NewOracleConnection() (*sql.DB, error) {
	user := os.Getenv("ORACLE_USER")
	password := os.Getenv("ORACLE_PASSWORD")
	connectString := os.Getenv("ORACLE_CONNECT_STRING")

	if user == "" || password == "" || connectString == "" {
		return nil, fmt.Errorf("Oracle environment variables are not configured")
	}

	dsn := fmt.Sprintf(
		`user="%s" password="%s" connectString="%s"`,
		user,
		password,
		connectString,
	)

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