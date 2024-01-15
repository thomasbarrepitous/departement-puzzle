package db_config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB(config DBConfig) (*sql.DB, error) {
	// Database connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.GetHost(),
		config.GetPort(),
		config.GetUsername(),
		config.GetPassword(),
		config.GetDatabase(),
	)
	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
