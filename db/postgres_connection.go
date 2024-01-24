package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB(config GameConfig) (*sql.DB, error) {
	// Database connection string
	// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.GetTypeDB(),
		config.GetUsername(),
		config.GetPassword(),
		config.GetHost(),
		config.GetPort(),
		config.GetDatabase(),
	)
	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
