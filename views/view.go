package views

import (
	"github.com/gofiber/fiber/v2"
	"tz.com/m/services"
)

type View struct {
	Ctx *fiber.Ctx
	PG  services.Database
	App *fiber.App
}
