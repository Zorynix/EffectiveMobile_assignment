package routes

import (
	"net/http"
)

func (router *Router) V1Routes() {

	http.HandleFunc("/v1/car-get", router.GetCarsRoute)
	http.HandleFunc("/v1/car-edit/", router.UpdateCarRoute)
	http.HandleFunc("/v1/car-add", router.AddCarRoute)
	http.HandleFunc("/v1/car-delete/", router.DeleteCarRoute)
}
