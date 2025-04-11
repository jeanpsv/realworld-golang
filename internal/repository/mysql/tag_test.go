package mysql_test

import (
	"errors"
	"testing"
	"time"

	"github.com/jeanpsv/realworld-golang/internal/repository/mysql"
	"github.com/jeanpsv/realworld-golang/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSelectShouldReturnList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	expectedTags := []*models.Tag{
		{
			ID:        1,
			Name:      "golang",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, {
			ID:        2,
			Name:      "unit test",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(expectedTags[0].ID, expectedTags[0].Name, expectedTags[0].CreatedAt, expectedTags[0].UpdatedAt).
		AddRow(expectedTags[1].ID, expectedTags[1].Name, expectedTags[1].CreatedAt, expectedTags[1].UpdatedAt)

	query := "SELECT id, name, created_at, updated_at FROM tags ORDER BY id"

	mock.ExpectQuery(query).WillReturnRows(rows)
	repo := mysql.NewTagRepository(db)
	actualTags, err := repo.Select()

	assert.Nil(t, err)

	for index, expected := range expectedTags {
		actual := actualTags[index]
		assert.EqualValues(t, expected, actual)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestSelectShouldReturnEmptyList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})

	query := "SELECT id, name, created_at, updated_at FROM tags ORDER BY id"

	mock.ExpectQuery(query).WillReturnRows(rows)
	repo := mysql.NewTagRepository(db)
	actualTags, err := repo.Select()

	assert.Nil(t, err)
	assert.Empty(t, actualTags)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestSelectShouldReturnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	expectedErrorMessage := "some error"

	query := "SELECT id, name, created_at, updated_at FROM tags ORDER BY id"

	mock.ExpectQuery(query).WillReturnError(errors.New(expectedErrorMessage))
	repo := mysql.NewTagRepository(db)
	actualTags, err := repo.Select()

	assert.Nil(t, actualTags)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErrorMessage, err.Error())
	assert.Nil(t, mock.ExpectationsWereMet())
}
