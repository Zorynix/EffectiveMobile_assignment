package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"

	"tz.com/m/services"
)

type RouterHead struct {
	PG   services.Database
	Addr *string
}

type Router struct {
	Router *fiber.App
	PG     services.Database
}

type Route struct {
	Group fiber.Router
	PG    services.Database
}

func Routes(addr *string) {
	postgres, err := services.NewPostgreSQL(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize PostgreSQL")
	}

	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	route := Router{Router: router, PG: postgres}

	route.V1Routes()

	log.Info().Msgf("Starting server on port %d...", 8000)
	if err := router.Listen(":8000"); err != nil {
		log.Fatal().Err(err).Msg("Cannot start HTTP server")
	}
}
