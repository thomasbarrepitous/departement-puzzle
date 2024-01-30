package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectDB connects to the database and returns a database connection
// It opens as many connections as the number of storages declated
// which is really bad but I wanted to create something modular similar
// to big projects with different DB
// I will probably refactor this into a singleton anyway in the future.
func ConnectDB(config PostgresConfig) (*sql.DB, error) {
	// Database connection string
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
