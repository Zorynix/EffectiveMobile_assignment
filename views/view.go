package views

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"tz.com/m/services"
)

type View struct {
	W  http.ResponseWriter
	R  *http.Request
	PG services.Database
}

func (view *View) respondWithJSON(data interface{}) {
	view.W.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(view.W).Encode(data); err != nil {
		view.handleError(err, http.StatusInternalServerError)
	}
}

func (view *View) handleError(err error, statusCode int) {
	log.Info().Err(err).Msg("")
	http.Error(view.W, http.StatusText(statusCode), statusCode)
}
