package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeanpsv/realworld-golang/services"
)

type TagHandler struct {
	service services.TagService
}

func NewTagHandler(router *mux.Router, s services.TagService) {
	handler := &TagHandler{
		service: s,
	}

	router.HandleFunc("/api/tags", handler.listTags).Methods("GET")
}

func (h *TagHandler) listTags(w http.ResponseWriter, r *http.Request) {
	tags, err := h.service.List()
	if err != nil {
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

	tagNames := []string{}
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}

	err = WriteJSON(w, http.StatusOK, Envelope{"tags": tagNames}, nil)
	if err != nil {
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
