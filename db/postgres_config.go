package db

import (
	"os"
)

type PostgresConfig struct {
	host     string
	port     string
	username string
	password string
	database string
	typeDB   string
}

// Abstract it at some point
func NewPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		username: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_NAME"),
		typeDB:   os.Getenv("DB_TYPE"),
	}
}

func (c *PostgresConfig) GetHost() string {
	return c.host
}

func (c *PostgresConfig) GetPort() string {
	return c.port
}

func (c *PostgresConfig) GetUsername() string {
	return c.username
}

func (c *PostgresConfig) GetPassword() string {
	return c.password
}

func (c *PostgresConfig) GetDatabase() string {
	return c.database
}

func (c *PostgresConfig) GetTypeDB() string {
	return c.typeDB
}
