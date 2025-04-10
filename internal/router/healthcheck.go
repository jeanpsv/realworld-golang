package router

import (
	"net/http"
)

const version = "1.0.0"

func (app *HttpApplication) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := Envelope{
		"status":      "available",
		"environment": app.Config.Env,
		"version":     version,
	}

	err := app.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.Logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
