package validation

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang-simple-boilerplate/exception"
	"golang-simple-boilerplate/model"
)

func BookValidation(Request model.BookRequest) (err error) {
	err = validation.ValidateStruct(&Request,
		validation.Field(&Request.Title, validation.Required.Error("NOT_BLANK")),
		validation.Field(&Request.Author, validation.Required.Error("NOT_BLANK")),
		validation.Field(&Request.Year, validation.Required.Error("NOT_BLANK")),
	)

	if err != nil {
		b, _ := json.Marshal(err)
		return exception.ValidationError{
			Message: string(b),
		}
	}

	return nil
}
