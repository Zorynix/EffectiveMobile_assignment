package views

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (view *View) GetCarsView(c *fiber.Ctx) error {

	filters := map[string]string{
		"reg_num":          c.Query("reg_num"),
		"mark":             c.Query("mark"),
		"model":            c.Query("model"),
		"year":             c.Query("year"),
		"owner_name":       c.Query("owner_name"),
		"owner_surname":    c.Query("owner_surname"),
		"owner_patronymic": c.Query("owner_patronymic"),
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	log.Info().Interface("filters", filters).Int("limit", limit).Int("offset", offset).Msg("GetCarsView called")

	data, err := view.PG.GetCars(filters, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("Error in GetCars")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	if data == nil {
		log.Error().Msg("Data is nil")
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	log.Info().Interface("data", data).Msg("Data")

	return c.JSON(*data)
}

func (view *View) UpdateCarView(c *fiber.Ctx) error {

	log.Info().Msg("UpdateCarView called")

	regNum := c.Query("regNum")
	var updates map[string]interface{}

	if err := c.BodyParser(&updates); err != nil {
		log.Error().Err(err).Msg("Failed to parse update body")
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	data, err := view.PG.UpdateCar(regNum, updates)
	if err != nil {
		log.Error().Err(err).Msg("Error in UpdateCars")
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	return c.JSON(*data)
}

func (view *View) AddCarView(c *fiber.Ctx) error {

	log.Info().Msg("AddCarView called")

	var req struct {
		RegNums []string `json:"regNums"`
	}

	if err := c.BodyParser(&req); err != nil {
		log.Error().Err(err).Msg("Failed to parse request body")
		return err
	}

	log.Debug().Interface("regNums", req.RegNums).Msg("Parsed registration numbers")

	data, err := view.PG.AddCar(req.RegNums)
	if err != nil {
		log.Error().Err(err).Msg("Error in AddCar")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return view.Ctx.JSON(*data)
}

func (view *View) DeleteCarView(c *fiber.Ctx) error {

	log.Info().Msg("DeleteCarView called")

	regNum := c.Query("regNum")
	data, err := view.PG.DeleteCar(regNum)
	if err != nil {
		log.Error().Err(err).Msg("Error in DeleteCar")
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	return c.JSON(*data)
}
