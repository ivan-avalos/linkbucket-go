/*
 *  validators.go
 *  Copyright (C) 2020  Iván Ávalos <ivan.avalos.diaz@hotmail.com>
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU Affero General Public License as
 *  published by the Free Software Foundation, either version 3 of the
 *  License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU Affero General Public License for more details.
 *
 *  You should have received a copy of the GNU Affero General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package setup

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
	"github.com/ivan-avalos/linkbucket-go/server/database"
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

// InitValidators initializes validators
func InitValidators(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Validator.(*CustomValidator).validator.RegisterTagNameFunc(tagName)
	e.Validator.(*CustomValidator).validator.RegisterValidation("unique", uniqueValidate)
	e.Validator.(*CustomValidator).validator.RegisterValidation("tags", tagsValidate)
}
