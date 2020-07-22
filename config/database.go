package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// DB containts information for current DB connection
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

// Build = build string of postgres gorm configuration
func (dbConfig *DBConfig) Build() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.DBName,
		dbConfig.Password,
	)
}

// DBConfigBuilder Create string of postgres gorm configuration
func DBConfigBuilder() string {
	if os.Getenv("DB_HOST")==nil{
		os.Setenv("DB_HOST")="127.0.0.1"
	}
	if os.Getenv("DB_PORT")==nil{
		os.Setenv("DB_PORT")="5432"
	}
	if os.Getenv("DB_USER")==nil{
		os.Setenv("DB_USER")="postgres"
	}
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASS"),
	}

	return dbConfig.Build()
}
