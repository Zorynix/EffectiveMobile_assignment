package services

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"tz.com/m/models"
)

func (PG *Postgresql) GetCars(w http.ResponseWriter, r *http.Request) (*models.Car, error) {

	log.Info().Msg("GetCars called")

	var data models.Car

	return &data, nil
}

func (PG *Postgresql) UpdateCar(w http.ResponseWriter, r *http.Request) (*models.Car, error) {

	log.Info().Msg("UpdateCar called")

	var data models.Car

	return &data, nil
}

func (PG *Postgresql) AddCar(w http.ResponseWriter, r *http.Request) (*models.Car, error) {
	log.Info().Msg("AddCar called")

	var data models.Car

	return &data, nil
}

func (PG *Postgresql) DeleteCar(w http.ResponseWriter, r *http.Request) (*models.Car, error) {

	log.Info().Msg("DeleteCar called")

	var data models.Car

	return &data, nil
}
