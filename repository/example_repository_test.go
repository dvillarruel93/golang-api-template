//go:build integration
// +build integration

package repository

import (
	"context"
	"empty-api-struct/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type ExampleRepositoryTestSuite struct {
	suite.Suite
	ExampleRepository *ExampleRepository
	db                *gorm.DB
	tearDownDB        func(r *require.Assertions)
}

func TestMockExampleRepository(t *testing.T) {
	suite.Run(t, new(ExampleRepositoryTestSuite))
}

func (s *ExampleRepositoryTestSuite) SetupTest() {
	sqlDB, gormDB, err := setupSqlDBAndGormDB()
	s.Require().NoError(err)

	s.db = gormDB
	s.tearDownDB = func(r *require.Assertions) {
		err := sqlDB.Close()
		r.NoError(err)
	}

	s.ExampleRepository = NewExampleRepository(s.db)
}

func (s *ExampleRepositoryTestSuite) TearDownTest() {
	s.tearDownDB(s.Require())
}

func (s *ExampleRepositoryTestSuite) TestAddAndFetchPerson() {
	personToSave := getPerson()

	// testing AddPerson func
	addResult, err := s.ExampleRepository.AddPerson(context.TODO(), personToSave)
	s.Require().Nil(err)

	// func to ensure test removes added person no matter if next validations fail or pass
	defer func() {
		s.Require().Nil(s.db.Delete(&addResult).Error)
	}()

	s.Require().NotNil(*addResult.ID)
	s.Require().Equal(personToSave.FirstName, addResult.FirstName)
	s.Require().Equal(personToSave.LastName, addResult.LastName)
	s.Require().Equal(personToSave.Address, addResult.Address)

	// testing FetchPersonByID func
	fetchResult, err := s.ExampleRepository.FetchPersonByID(context.TODO(), *addResult.ID)
	s.Require().Nil(err)
	s.Require().Equal(*addResult.ID, *fetchResult.ID)
	s.Require().Equal(addResult.FirstName, fetchResult.FirstName)
	s.Require().Equal(addResult.LastName, fetchResult.LastName)
	s.Require().Equal(addResult.Address, fetchResult.Address)

}

func getPerson() models.Person {
	return models.Person{
		FirstName: "a name",
		LastName:  "a last name",
		Address:   "an address",
	}
}
