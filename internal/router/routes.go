package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *HttpApplication) Routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/v1/healthcheck", app.healthcheckHandler).Methods("GET")
	router.HandleFunc("/api/tags", app.listTagsHandler).Methods("GET")

	return router
}
