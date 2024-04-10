package routes

import (
	"net/http"

	"tz.com/m/views"
)

func (router *Router) GetCarsRoute(w http.ResponseWriter, r *http.Request) {
	view := views.View{W: w, R: r, PG: router.PG}
	view.GetCarsView()
}

func (router *Router) UpdateCarRoute(w http.ResponseWriter, r *http.Request) {
	view := views.View{W: w, R: r, PG: router.PG}
	view.UpdateCarView()
}

func (router *Router) AddCarRoute(w http.ResponseWriter, r *http.Request) {
	view := views.View{W: w, R: r, PG: router.PG}
	view.AddCarView()
}

func (router *Router) DeleteCarRoute(w http.ResponseWriter, r *http.Request) {
	view := views.View{W: w, R: r, PG: router.PG}
	view.DeleteCarView()
}
