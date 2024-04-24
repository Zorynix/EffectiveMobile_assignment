package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"tz.com/m/views"
)

func (route *Route) GetCarsRoute() {
	route.Group.Get("/info", func(c *fiber.Ctx) error {
		view := views.View{PG: route.PG}
		log.Info().Interface("ctx", c).Msg("CTX")
		return view.GetCarsView(c)
	})
}

func (route *Route) UpdateCarRoute() {
	route.Group.Put("/car-edit", func(c *fiber.Ctx) error {
		view := views.View{PG: route.PG}
		return view.UpdateCarView(c)
	})
}

func (route *Route) AddCarRoute() {
	route.Group.Post("/car-add", func(c *fiber.Ctx) error {
		view := views.View{PG: route.PG}
		return view.AddCarView(c)
	})
}

func (route *Route) DeleteCarRoute() {
	route.Group.Delete("/car-delete", func(c *fiber.Ctx) error {
		view := views.View{PG: route.PG}
		return view.DeleteCarView(c)
	})
}
