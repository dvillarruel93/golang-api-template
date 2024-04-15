package models

import (
	"empty-api-struct/api_error"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var (
	validate *validator.Validate // use a singleton to be thread safe and cache information per docs
)

func init() {
	validate = validator.New()
}

// Validate runs through the tags and validates all fields.
func Validate(model interface{}) error {
	if err := validate.Struct(model); err != nil {
		e, ok := err.(validator.ValidationErrors)
		if ok {
			return api_error.New(http.StatusBadRequest, "error validating model").WithInternal(e)
		}
		return err
	}
	return nil
}
