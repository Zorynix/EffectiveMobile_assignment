package routes

import (
	"github.com/gofiber/fiber/v2"
	"tz.com/m/views"
)

func (route *Route) GetCarsRoute() {
	route.Group.Get("/car-get", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, PG: route.PG}
		return view.GetCarsView()
	})
}

func (route *Route) UpdateCarRoute() {
	route.Group.Get("/car-edit", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, PG: route.PG}
		return view.GetCarsView()
	})
}

func (route *Route) AddCarRoute() {
	route.Group.Get("/car-add", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, PG: route.PG}
		return view.GetCarsView()
	})
}

func (route *Route) DeleteCarRoute() {
	route.Group.Get("/car-delete", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, PG: route.PG}
		return view.GetCarsView()
	})
}
