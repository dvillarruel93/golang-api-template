package interfaces

import (
	"context"
	"empty-api-struct/models"
)

type ExampleUsecase interface {
	Test() string
	AddPerson(ctx context.Context, person models.Person) (models.Person, error)
	FetchPersonByID(ctx context.Context, id string) (models.Person, error)
}
