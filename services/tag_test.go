package services

import (
	"errors"
	"testing"
	"time"

	"github.com/jeanpsv/realworld-golang/mocks"
	"github.com/jeanpsv/realworld-golang/models"
	"github.com/stretchr/testify/assert"
)

func TestSelectShouldReturnList(t *testing.T) {
	expectedTags := []*models.Tag{
		{
			ID:        1,
			Name:      "testing",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockedTagRepository := mocks.NewTagRepository(t)
	mockedTagRepository.On("Select").Return(expectedTags, nil)

	tagService := NewTagService(mockedTagRepository)
	actualTags, err := tagService.List()
	assert.Nil(t, err)

	for index, expected := range expectedTags {
		actual := actualTags[index]
		assert.EqualValues(t, expected, actual)
	}
}

func TestSelectShouldReturnEmptyList(t *testing.T) {
	mockedTagRepository := mocks.NewTagRepository(t)
	mockedTagRepository.On("Select").Return([]*models.Tag{}, nil)

	tagService := NewTagService(mockedTagRepository)
	actualTags, err := tagService.List()
	assert.Nil(t, err)
	assert.Empty(t, actualTags)
}

func TestSelectShouldReturnError(t *testing.T) {
	expectedErrorMessage := "some error"
	mockedTagRepository := mocks.NewTagRepository(t)
	mockedTagRepository.On("Select").Return(nil, errors.New(expectedErrorMessage))

	tagService := NewTagService(mockedTagRepository)
	actualTags, err := tagService.List()

	assert.Nil(t, actualTags)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErrorMessage, err.Error())
}
