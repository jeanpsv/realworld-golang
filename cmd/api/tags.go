package main

import "net/http"

func (app *application) listTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := app.models.Tags.List()
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

	tagNames := []string{}
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"tags": tagNames}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
