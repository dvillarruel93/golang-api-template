package rest

import (
	"empty-api-struct/api_error"
	"empty-api-struct/helper/appcontext"
	"empty-api-struct/interfaces"
	"empty-api-struct/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ExampleHandler struct {
	ExampleUsecase interfaces.ExampleUsecase
	AuthMW         *echo.MiddlewareFunc
}

func NewExampleHandler(e *echo.Echo, exampleUsecase interfaces.ExampleUsecase, authMW *echo.MiddlewareFunc) {
	handler := &ExampleHandler{
		ExampleUsecase: exampleUsecase,
		AuthMW:         authMW,
	}

	e.GET("/v1/test", handler.Test, *handler.AuthMW)
	e.POST("/v1/person", handler.AddPerson, *handler.AuthMW)
	e.GET("/v1/person/:personID", handler.FetchPersonByID, *handler.AuthMW)
}

func (handler *ExampleHandler) Test(c echo.Context) error {
	response := fmt.Sprintf("%s - %s", "handler", handler.ExampleUsecase.Test())
	return c.JSON(http.StatusOK, response)
}

func (handler *ExampleHandler) AddPerson(c echo.Context) error {
	person := models.Person{}
	if err := c.Bind(&person); err != nil {
		return api_error.New(http.StatusBadRequest, "error binding body").WithInternal(err)
	}

	if err := models.Validate(person); err != nil {
		return api_error.New(http.StatusBadRequest, "error validating body").WithInternal(err)
	}

	response, err := handler.ExampleUsecase.AddPerson(appcontext.EchoContextToContext(c), person)
	if err != nil {
		return api_error.New(http.StatusInternalServerError, "error adding person").WithInternal(err)
	}

	return c.JSON(http.StatusOK, response)
}

func (handler *ExampleHandler) FetchPersonByID(c echo.Context) error {
	id := c.Param("personID")
	if len(id) == 0 {
		return api_error.New(http.StatusBadRequest, "personID param could not be empty")
	}

	response, err := handler.ExampleUsecase.FetchPersonByID(appcontext.EchoContextToContext(c), id)
	if err != nil {
		return api_error.New(http.StatusInternalServerError, "error adding person").WithInternal(err)
	}

	return c.JSON(http.StatusOK, response)
}
