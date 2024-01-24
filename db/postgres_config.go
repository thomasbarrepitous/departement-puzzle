package db

import (
	"os"
)

type GameConfig struct {
	host     string
	port     string
	username string
	password string
	database string
	typeDB   string
}

func NewGameConfig() *GameConfig {
	return &GameConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		username: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_NAME"),
		typeDB:   os.Getenv("DB_TYPE"),
	}
}

func (c *GameConfig) GetHost() string {
	return c.host
}

func (c *GameConfig) GetPort() string {
	return c.port
}

func (c *GameConfig) GetUsername() string {
	return c.username
}

func (c *GameConfig) GetPassword() string {
	return c.password
}

func (c *GameConfig) GetDatabase() string {
	return c.database
}

func (c *GameConfig) GetTypeDB() string {
	return c.typeDB
}
