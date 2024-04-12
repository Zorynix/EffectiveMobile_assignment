package routes

import (
	"github.com/gofiber/fiber/v2"
	"tz.com/m/views"
)

func (route *Route) GetCarsRoute() {
	route.Group.Get("/info", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, PG: route.PG}
		return view.GetCarsView()
	})
}

func (route *Route) UpdateCarRoute() {
	route.Group.Put("/car-edit", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, PG: route.PG}
		return view.UpdateCarView()
	})
}

func (route *Route) AddCarRoute() {
	route.Group.Post("/car-add", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, PG: route.PG}
		return view.AddCarView()
	})
}

func (route *Route) DeleteCarRoute() {
	route.Group.Delete("/car-delete", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, PG: route.PG}
		return view.DeleteCarView()
	})
}
