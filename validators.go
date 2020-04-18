package main

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
	"github.com/ivan-avalos/linkbucket/database"
	"github.com/labstack/echo"
)

// CustomValidator is a custom validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates using a CustomValidator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func tagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

func uniqueValidate(fld validator.FieldLevel) bool {
	// table.field
	param := strings.Split(fld.Param(), ".")
	table := param[0]
	field := param[1]
	count := 0
	database.DB().Table(table).
		Where(field+" = ?", fld.Field().String()).Count(&count)
	return count == 0
}

func tagsValidate(fld validator.FieldLevel) bool {
	ok, err := regexp.MatchString("^([^\\,\\n\\t]{1,30}\\,){0,30}([^\\,\\n\\t]{1,30})?$", fld.Field().String())
	if err != nil {
		return false
	}
	return ok
}

func initValidators(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Validator.(*CustomValidator).validator.RegisterTagNameFunc(tagName)
	e.Validator.(*CustomValidator).validator.RegisterValidation("unique", uniqueValidate)
	e.Validator.(*CustomValidator).validator.RegisterValidation("tags", tagsValidate)
}
