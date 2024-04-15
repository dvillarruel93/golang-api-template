package usecase

import (
	"context"
	"empty-api-struct/interfaces"
	"empty-api-struct/models"
	"fmt"
)

type ExampleUsecase struct {
	ExampleRepository interfaces.ExampleRepository
}

func NewExampleUsecase(exampleRepository interfaces.ExampleRepository) *ExampleUsecase {
	return &ExampleUsecase{
		ExampleRepository: exampleRepository,
	}
}

func (eu *ExampleUsecase) Test() string {
	str := "example usecase"
	return fmt.Sprintf("%s - %s", str, eu.ExampleRepository.Test())
}

func (eu *ExampleUsecase) AddPerson(ctx context.Context, person models.Person) (models.Person, error) {
	return eu.ExampleRepository.AddPerson(ctx, person)
}

func (eu *ExampleUsecase) FetchPersonByID(ctx context.Context, id string) (models.Person, error) {
	return eu.ExampleRepository.FetchPersonByID(ctx, id)
}
