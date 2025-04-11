package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/jeanpsv/realworld-golang/mocks"
	"github.com/jeanpsv/realworld-golang/models"
	"github.com/stretchr/testify/assert"
)

func TestListHandler(t *testing.T) {
	expectedTags := []*models.Tag{
		{
			ID:        1,
			Name:      "testing",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockedTagService := mocks.NewTagUseCase(t)
	mockedTagService.On("List").Return(expectedTags, nil)

	req, err := http.NewRequest("GET", "/api/tags", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	NewTagHandler(router, mockedTagService)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	print(rr.Body.String())
}
