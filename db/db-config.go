package db

import (
	"os"
)

type DBConfig struct {
	host     string
	port     string
	username string
	password string
	database string
	typeDB   string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		username: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_NAME"),
		typeDB:   os.Getenv("DB_TYPE"),
	}
}

func (c *DBConfig) GetHost() string {
	return c.host
}

func (c *DBConfig) GetPort() string {
	return c.port
}

func (c *DBConfig) GetUsername() string {
	return c.username
}

func (c *DBConfig) GetPassword() string {
	return c.password
}

func (c *DBConfig) GetDatabase() string {
	return c.database
}

func (c *DBConfig) GetTypeDB() string {
	return c.typeDB
}
