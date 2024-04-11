package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (view *View) GetCarsView() error {

	log.Info().Msg("GetCarsView called")

	data, err := view.PG.GetCars(view.Ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error in GetCars")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return view.Ctx.JSON(*data)
}

func (view *View) UpdateCarView() error {

	log.Info().Msg("UpdateCarView called")

	data, err := view.PG.UpdateCar(view.Ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error in UpdateCars")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return view.Ctx.JSON(*data)
}

func (view *View) AddCarView() error {

	log.Info().Msg("AddCarView called")

	data, err := view.PG.AddCar(view.Ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error in AddCar")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return view.Ctx.JSON(*data)
}

func (view *View) DeleteCarView() error {

	log.Info().Msg("DeleteCarView called")

	data, err := view.PG.DeleteCar(view.Ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error in DeleteCar")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return view.Ctx.JSON(*data)
}
