package db

import "github.com/spf13/viper"

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
		host:     viper.Get("DB_HOST").(string),
		port:     viper.Get("DB_PORT").(int),
		username: viper.Get("DB_USERNAME").(string),
		password: viper.Get("DB_PASSWORD").(string),
		database: viper.Get("DB_NAME").(string),
		typeDB:   viper.Get("DB_TYPE").(string),
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
