package rest

import (
	"encoding/json"
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
	tag1 := models.Tag{
		ID:        1,
		Name:      "testing",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tag2 := models.Tag{
		ID:        2,
		Name:      "golang",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tags := []*models.Tag{&tag1, &tag2}

	mockedTagService := mocks.NewTagUseCase(t)
	mockedTagService.On("List").Return(tags, nil)

	req, err := http.NewRequest("GET", "/api/tags", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	NewTagHandler(router, mockedTagService)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var body = map[string][]string{}
	err = json.Unmarshal(rr.Body.Bytes(), &body)
	assert.Nil(t, err)
	assert.NotNil(t, body["tags"])
	expectedTags := []string{tag1.Name, tag2.Name}
	assert.Exactly(t, expectedTags, body["tags"])
}
