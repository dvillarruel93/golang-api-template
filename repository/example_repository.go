package repository

import (
	"context"
	"empty-api-struct/api_error"
	"empty-api-struct/models"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

type ExampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) *ExampleRepository {
	return &ExampleRepository{
		db: db,
	}
}

func (r *ExampleRepository) Test() string {
	return "example repo"
}

func (r *ExampleRepository) AddPerson(ctx context.Context, person models.Person) (models.Person, error) {
	if err := r.db.WithContext(ctx).Create(&person).Error; err != nil {
		return models.Person{}, api_error.New(http.StatusInternalServerError, "error creating person").
			WithInternal(err)
	}

	return person, nil
}

func (r *ExampleRepository) FetchPersonByID(ctx context.Context, id string) (models.Person, error) {
	result := models.Person{}
	if err := r.db.WithContext(ctx).Model(&models.Person{}).Where("id = ?", id).Find(&result).Error; err != nil {
		return models.Person{}, api_error.New(http.StatusInternalServerError,
			fmt.Sprintf("error fetching person by id %s", id)).
			WithInternal(err)
	}

	return result, nil
}
