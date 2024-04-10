package services

import (
	"context"
	"errors"
	"net/http"
	"os"

	"tz.com/m/models"
	"tz.com/m/utils"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Ping(ctx context.Context) error

	GetCars(w http.ResponseWriter, r *http.Request) (*[]models.Car, error)
	UpdateCar(w http.ResponseWriter, r *http.Request) (*models.Car, error)
	AddCar(w http.ResponseWriter, r *http.Request) (*models.Car, error)
	DeleteCar(w http.ResponseWriter, r *http.Request) (*models.Car, error)
}

type Postgresql struct {
	DB *gorm.DB
}

// NewPostgreSQL creates and returns a new Postgresql instance
// This function initializes a PostgreSQL database connection using the DSN environment variable
// It sets the search path to 'tz' and automatically migrates the database schemas for Car model
// Returns a pointer to a Postgresql struct or an error if the connection or migration fails
func NewPostgreSQL(ctx context.Context) (Database, error) {
	utils.LoadEnv()

	DSN := os.Getenv("DSN")
	if DSN == "" {
		return nil, errors.New("DSN is not set")
	}

	conn, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn = conn.Debug()

	conn.Exec("SET search_path TO tz")

	err = conn.AutoMigrate(&models.Car{}, &models.People{})
	if err != nil {
		return nil, err
	}

	return &Postgresql{DB: conn}, nil
}

// Ping checks the connection to the PostgreSQL database
// It verifies that the database is accessible and responding to queries
// Returns an error if the database is unreachable or not responding
func (pg *Postgresql) Ping(ctx context.Context) error {
	DB, err := pg.DB.DB()
	if err != nil {
		log.Fatal().Interface("unable to create postgresql connection pool: %v", err).Msg("")
		return err
	}
	return DB.PingContext(ctx)
}

// Close terminates the PostgreSQL database connection
// It safely closes the connection pool, freeing up resources
// Logs a fatal error if closing the connection pool fails
func (pg *Postgresql) Close() {

	DB, err := pg.DB.DB()
	if err != nil {
		log.Fatal().Interface("unable to create postgresql connection pool: %v", err).Msg("")
	}
	DB.Close()
}
