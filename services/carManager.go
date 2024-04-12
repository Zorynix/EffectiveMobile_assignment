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
	log.Debug().Msg("Starting GetCars method")
	log.Info().Msg("GetCars called")

	var cars []models.Car
	query := PG.DB.Preload("Owner")
	log.Debug().Msg("Preloaded Owner")

	modelFilters := []string{"reg_num", "mark", "model", "year", "owner_name", "owner_surname", "owner_patronymic"}
	for _, filter := range modelFilters {
		if value := c.Query(filter); value != "" {
			query = query.Where(filter+" = ?", value)
			log.Debug().Str("filter", filter).Str("value", value).Msg("Applying filter")
		}
	}

	var limit, offset int
	if v := c.Query("limit"); v != "" {
		limit, _ = strconv.Atoi(v)
		query = query.Limit(limit)
		log.Debug().Int("limit", limit).Msg("Setting limit")
	}
	if v := c.Query("offset"); v != "" {
		offset, _ = strconv.Atoi(v)
		query = query.Offset(offset)
		log.Debug().Int("offset", offset).Msg("Setting offset")
	}

	if err := query.Find(&cars).Error; err != nil {
		log.Error().Err(err).Msg("Failed to fetch cars")
		return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info().Msg("Successfully fetched cars")
	return &cars, nil
}

func (PG *Postgresql) UpdateCar(c *fiber.Ctx) (*models.Car, error) {
	log.Debug().Msg("Starting UpdateCar method")
	log.Info().Msg("UpdateCar called")

	var car models.Car
	regNum := c.Query("regNum")
	log.Debug().Str("regNum", regNum).Msg("Received regNum for updating")

	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		log.Error().Err(err).Msg("Failed to parse update body")
		return nil, err
	}
	log.Debug().Interface("updates", updates).Msg("Parsed updates")

	if result := PG.DB.Model(&car).Where("reg_num = ?", regNum).Updates(updates); result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to update car")
		return nil, result.Error
	}

	if result := PG.DB.Where("reg_num = ?", regNum).First(&car); result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to retrieve updated car")
		return nil, result.Error
	}

	log.Info().Msg("Successfully updated car")
	return &car, nil
}

func fetchCarInfo(regNum string) (*models.Car, error) {
	log.Debug().Str("regNum", regNum).Msg("Starting fetchCarInfo")
	utils.LoadEnv()

	var car models.Car

	link := os.Getenv("URL")
	log.Debug().Str("URL", link).Msg("Loaded environment URL")

	url := fmt.Sprintf(link, regNum)
	log.Debug().Str("Formatted URL", url).Msg("Formatted URL with regNum")

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create HTTP request")
		return &car, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("HTTP request failed")
		return &car, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().Int("StatusCode", resp.StatusCode).Msg("Received non-OK HTTP status")
		return &car, fmt.Errorf("failed to fetch car info, status code: %d", resp.StatusCode)
	}

	if err = json.NewDecoder(resp.Body).Decode(&car); err != nil {
		log.Error().Err(err).Msg("Failed to decode car info")
		return &car, err
	}

	log.Info().Msg("Successfully fetched car info")
	return &car, nil
}

func (PG *Postgresql) AddCar(c *fiber.Ctx) (*models.Car, error) {
	log.Debug().Msg("Starting AddCar method")
	log.Info().Msg("AddCar called")

	var car models.Car

	var req struct {
		RegNums []string `json:"regNums"`
	}

	if err := c.BodyParser(&req); err != nil {
		log.Error().Err(err).Msg("Failed to parse request body")
		return nil, err
	}
	log.Debug().Interface("regNums", req.RegNums).Msg("Parsed registration numbers")

	for _, regNum := range req.RegNums {
		carInfo, err := fetchCarInfo(regNum)
		if err != nil {
			log.Error().Str("regNum", regNum).Err(err).Msg("Failed to fetch car info")
			continue
		}

		if result := PG.DB.Create(&carInfo); result.Error != nil {
			log.Error().Err(result.Error).Msg("Failed to add car to database")
			return nil, result.Error
		}

		return &car, c.JSON(&carInfo)
	}

	log.Warn().Msg("No cars were added")
	return &car, errors.New("no cars were added")
}

func (PG *Postgresql) DeleteCar(c *fiber.Ctx) (*models.Car, error) {
	log.Debug().Msg("Starting DeleteCar method")
	log.Info().Msg("DeleteCar called")

	var car models.Car

	regNum := c.Query("regNum")
	log.Debug().Str("regNum", regNum).Msg("Received regNum for deletion")

	if result := PG.DB.Where("reg_num = ?", regNum).Delete(&car); result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to delete car")
		return nil, result.Error
	}

	log.Info().Msg("Successfully deleted car")
	return &car, nil
}
