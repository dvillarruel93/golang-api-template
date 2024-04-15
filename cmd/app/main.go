package main

import (
	"empty-api-struct/delivery/middleware"
	"empty-api-struct/delivery/rest"
	"empty-api-struct/repository"
	"empty-api-struct/usecase"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	mysqlDB, err := repository.SetupDB()
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err)
	}

	e := echo.New()
	e.Debug = true
	e.HideBanner = true

	exampleRepo := repository.NewExampleRepository(mysqlDB)
	exampleUsecase := usecase.NewExampleUsecase(exampleRepo)
	authFnc := middleware.AuthMW()
	rest.NewExampleHandler(e, exampleUsecase, &authFnc)

	e.Logger.Fatal(e.Start(":8080"))
}
