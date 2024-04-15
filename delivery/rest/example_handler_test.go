//go:build unit
// +build unit

package rest

import (
	"bytes"
	"empty-api-struct/api_error"
	helper "empty-api-struct/helper/pointer"
	"empty-api-struct/mocks"
	"empty-api-struct/models"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

var errMock = errors.New("mock error")

type ExampleHandlerTestSuite struct {
	suite.Suite
	Handler        ExampleHandler
	ExampleUsecase *mocks.ExampleUsecase
}

func TestMockExampleHandler(t *testing.T) {
	suite.Run(t, new(ExampleHandlerTestSuite))
}

func (s *ExampleHandlerTestSuite) SetupTest() {
	s.ExampleUsecase = new(mocks.ExampleUsecase)
	s.Handler = ExampleHandler{
		ExampleUsecase: s.ExampleUsecase,
		AuthMW:         nil,
	}
}

func (s *ExampleHandlerTestSuite) TestAddPerson() {
	personWithAnEmptyField := getPerson("a random id")
	personWithAnEmptyField.FirstName = ""
	var tests = []struct {
		name      string
		person    models.Person
		wantErr   bool
		errString string
	}{
		{
			name:      "error validating person body",
			person:    personWithAnEmptyField,
			wantErr:   true,
			errString: "error validating body",
		},
		{
			name:      "error validating person body",
			person:    getPerson("a wrong id"),
			wantErr:   true,
			errString: "error adding person",
		},
		{
			name:      "error validating person body",
			person:    getPerson("an ok id"),
			wantErr:   false,
			errString: "",
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// set echo context
			bodyBytes, err := json.Marshal(test.person)
			s.Require().NoError(err)
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bodyBytes))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			recorder := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(request, recorder)

			// prepare mock result
			var mockedError error
			if test.wantErr {
				mockedError = errMock
			}
			s.ExampleUsecase.
				On("AddPerson", mock.Anything, test.person).
				Return(test.person, mockedError)

			err = s.Handler.AddPerson(ctx)

			if test.wantErr {
				s.Require().NotNil(err)
				s.Require().Contains(err.Error(), test.errString)
				return
			}

			s.Require().Nil(err)
			var person models.Person
			err = json.Unmarshal(recorder.Body.Bytes(), &person)
			s.Require().NoError(err)
			s.Require().Equal(test.person, person)
		})
	}
}

func (s *ExampleHandlerTestSuite) TestFetchPersonByID() {
	var tests = []struct {
		name     string
		personID string
		wantErr  bool
		err      error
	}{
		{
			name:     "error empty personID param",
			personID: "",
			wantErr:  true,
			err:      api_error.New(http.StatusBadRequest, "personID param could not be empty"),
		},
		{
			name:     "error fetching person",
			personID: "wrong id",
			wantErr:  true,
			err:      api_error.New(http.StatusInternalServerError, "error adding person").WithInternal(errMock),
		},
		{
			name:     "ok",
			personID: "ok id",
			wantErr:  false,
			err:      nil,
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// set echo context
			request := httptest.NewRequest(http.MethodGet, "/", nil)
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			recorder := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(request, recorder)
			ctx.SetParamNames("personID")
			ctx.SetParamValues(test.personID)

			// prepare mock result
			mockedPerson := getPerson(test.personID)
			var mockedError error
			if test.wantErr {
				mockedError = errMock
			}
			s.ExampleUsecase.
				On("FetchPersonByID", mock.Anything, test.personID).
				Return(mockedPerson, mockedError)

			err := s.Handler.FetchPersonByID(ctx)

			if test.wantErr {
				s.Require().NotNil(err)
				s.Require().EqualError(test.err, err.Error())
				return
			}

			s.Require().Nil(err)
			var person models.Person
			err = json.Unmarshal(recorder.Body.Bytes(), &person)
			s.Require().NoError(err)
			s.Require().Equal(mockedPerson, person)
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
