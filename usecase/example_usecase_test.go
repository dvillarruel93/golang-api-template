//go:build unit
// +build unit

package usecase

import (
	"context"
	helper "empty-api-struct/helper/pointer"
	"empty-api-struct/mocks"
	"empty-api-struct/models"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

var errMock = errors.New("mock error")

type MockExampleUsecaseTestSuite struct {
	suite.Suite
	ExampleUsecase    ExampleUsecase
	ExampleRepository *mocks.ExampleRepository
}

func TestMockExampleUsecase(t *testing.T) {
	suite.Run(t, new(MockExampleUsecaseTestSuite))
}

func (s *MockExampleUsecaseTestSuite) SetupTest() {
	s.ExampleRepository = new(mocks.ExampleRepository)
	s.ExampleUsecase = ExampleUsecase{
		ExampleRepository: s.ExampleRepository,
	}
}

func (s *MockExampleUsecaseTestSuite) TestAddPerson() {
	var tests = []struct {
		name    string
		person  models.Person
		wantErr bool
		err     error
	}{
		{
			name:    "error adding person",
			wantErr: true,
			person:  getPerson("a wrong test id"),
			err:     errMock,
		},
		{
			name:    "ok",
			wantErr: false,
			person:  getPerson("a test id"),
			err:     nil,
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// prepare mock result
			s.ExampleRepository.
				On("AddPerson", mock.Anything, test.person).
				Return(test.person, test.err)

			result, err := s.ExampleUsecase.AddPerson(context.TODO(), test.person)

			if test.wantErr {
				s.Require().NotNil(err)
				s.Require().Equal(test.err, err)
				return
			}

			s.Require().Nil(err)
			s.Require().Equal(test.person, result)
		})
	}
}

func (s *MockExampleUsecaseTestSuite) TestFetchPersonByID() {
	var tests = []struct {
		name    string
		id      string
		wantErr bool
		err     error
	}{
		{
			name:    "error fetching person",
			wantErr: true,
			id:      "wrong id",
			err:     errMock,
		},
		{
			name:    "ok",
			wantErr: false,
			id:      "ok id",
			err:     nil,
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// prepare mock result
			mockedPerson := getPerson(test.id)
			s.ExampleRepository.
				On("FetchPersonByID", mock.Anything, test.id).
				Return(mockedPerson, test.err)

			result, err := s.ExampleUsecase.FetchPersonByID(context.TODO(), test.id)

			if test.wantErr {
				s.Require().NotNil(err)
				s.Require().Equal(test.err, err)
				return
			}

			s.Require().Nil(err)
			s.Require().Equal(mockedPerson, result)
		})
	}
}

func getPerson(id string) models.Person {
	return models.Person{
		ModelBase: models.ModelBase{
			ID: helper.ToStringPtr(id),
		},
		FirstName: "a name",
		LastName:  "a last name",
		Address:   "an address",
	}
}
