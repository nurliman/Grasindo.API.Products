package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// DB containts information for current DB connection
var (
	DB  *gorm.DB
	err error
)

// DBConfig represents db configuration
type dbConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

// build = build string of postgres gorm configuration
func build() string {
	if os.Getenv("DB_HOST") == "" {
		os.Setenv("DB_HOST", "127.0.0.1")
	}
	if os.Getenv("DB_PORT") == "" {
		os.Setenv("DB_PORT", "5432")
	}
	if os.Getenv("DB_USER") == "" {
		os.Setenv("DB_USER", "postgres")
	}
	db := &dbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASS"),
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		db.Host,
		db.Port,
		db.User,
		db.DBName,
		db.Password,
	)
}

// DBInit initialize database
func DBInit() {
	DB, err = gorm.Open("postgres", build())
	if err != nil {
		panic(err)
	}

}
