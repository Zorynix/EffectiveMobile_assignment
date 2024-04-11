package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"tz.com/m/models"
	"tz.com/m/utils"
)

func (PG *Postgresql) GetCars(c *fiber.Ctx) (*[]models.Car, error) {

	log.Info().Msg("GetCars called")

	var cars []models.Car
	query := PG.DB

	modelFilters := []string{"reg_num", "mark", "model", "year", "owner_name", "owner_surname", "owner_patronymic"}
	for _, filter := range modelFilters {
		if value := c.Query(filter); value != "" {
			query = query.Where(filter+" = ?", value)
		}
	}

	var limit, offset int
	if v := c.Query("limit"); v != "" {
		limit, _ = strconv.Atoi(v)
		query = query.Limit(limit)
	}
	if v := c.Query("offset"); v != "" {
		offset, _ = strconv.Atoi(v)
		query = query.Offset(offset)
	}

	if err := query.Find(&cars).Error; err != nil {
		log.Error().Err(err).Msg("Failed to fetch cars")
		return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return &cars, nil
}

func (PG *Postgresql) UpdateCar(c *fiber.Ctx) (*models.Car, error) {

	log.Info().Msg("UpdateCar called")

	regNum := c.Query("regNum")
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return nil, err
	}
	if result := PG.DB.Model(&models.Car{}).Where("reg_num = ?", regNum).Updates(updates); result.Error != nil {
		return nil, result.Error
	}

	var car models.Car
	if result := PG.DB.Where("reg_num = ?", regNum).First(&car); result.Error != nil {
		return nil, result.Error
	}
	return &car, nil
}

func fetchCarInfo(regNum string) (*models.Car, error) {

	utils.LoadEnv()

	link := os.Getenv("URL")

	url := fmt.Sprintf(link, regNum)

	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &models.Car{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return &models.Car{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &models.Car{}, fmt.Errorf("failed to fetch car info, status code: %d", resp.StatusCode)
	}

	var car models.Car
	if err = json.NewDecoder(resp.Body).Decode(&car); err != nil {
		return &models.Car{}, err
	}

	return &car, nil
}

func (PG *Postgresql) AddCar(c *fiber.Ctx) (*models.Car, error) {

	log.Info().Msg("AddCar called")

	var req struct {
		RegNums []string `json:"regNums"`
	}

	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}

	for _, regNum := range req.RegNums {

		carInfo, err := fetchCarInfo(regNum)
		if err != nil {
			continue
		}

		if result := PG.DB.Create(&carInfo); result.Error != nil {
			return nil, result.Error
		}

		return &models.Car{}, c.JSON(&carInfo)
	}

	return &models.Car{}, errors.New("no cars were added")
}

func (PG *Postgresql) DeleteCar(c *fiber.Ctx) (*models.Car, error) {

	log.Info().Msg("DeleteCar called")

	regNum := c.Query("regNum")
	if result := PG.DB.Where("reg_num = ?", regNum).Delete(&models.Car{}); result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to delete car")
		return nil, result.Error
	}
	return &models.Car{}, nil
}
