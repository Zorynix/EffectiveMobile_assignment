package views

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (view *View) GetCarsView() error {

	log.Info().Msg("GetCarsView called")

	data, err := view.PG.GetCars(view.W, view.R)
	if err != nil {
		log.Error().Err(err).Msg("Error in GetCars")
		view.handleError(err, http.StatusBadGateway)
		return err
	}

	view.respondWithJSON(data)
	return nil
}

func (view *View) UpdateCarView() error {

	log.Info().Msg("UpdateCarView called")

	data, err := view.PG.UpdateCar(view.W, view.R)
	if err != nil {
		log.Error().Err(err).Msg("Error in UpdateCars")
		view.handleError(err, http.StatusBadGateway)
		return err
	}

	view.respondWithJSON(data)
	return nil
}

func (view *View) AddCarView() error {

	log.Info().Msg("AddCarView called")

	data, err := view.PG.AddCar(view.W, view.R)
	if err != nil {
		log.Error().Err(err).Msg("Error in AddCar")
		view.handleError(err, http.StatusBadGateway)
		return err
	}

	view.respondWithJSON(data)
	return nil
}

func (view *View) DeleteCarView() error {

	log.Info().Msg("DeleteCarView called")

	data, err := view.PG.DeleteCar(view.W, view.R)
	if err != nil {
		log.Error().Err(err).Msg("Error in DeleteCar")
		view.handleError(err, http.StatusBadGateway)
		return err
	}

	view.respondWithJSON(data)
	return nil
}
