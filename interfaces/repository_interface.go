package interfaces

import (
	"context"
	"empty-api-struct/models"
)

type ExampleRepository interface {
	Test() string
	AddPerson(ctx context.Context, person models.Person) (models.Person, error)
	FetchPersonByID(ctx context.Context, id string) (models.Person, error)
}
