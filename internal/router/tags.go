package router

import "net/http"

func (app *HttpApplication) listTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := app.Models.Tags.List()
	if err != nil {
		app.Logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

	tagNames := []string{}
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}

	err = app.WriteJSON(w, http.StatusOK, Envelope{"tags": tagNames}, nil)
	if err != nil {
		app.Logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
