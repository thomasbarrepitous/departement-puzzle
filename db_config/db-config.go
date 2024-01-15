package db_config

import "os"

type DBConfig struct {
	host     string
	port     string
	username string
	password string
	database string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		host:     os.Getenv("POSTGRES_HOST"),
		port:     os.Getenv("POSTGRES_PORT"),
		username: os.Getenv("POSTGRES_USERNAME"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		database: os.Getenv("POSTGRES_DATABASE"),
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
