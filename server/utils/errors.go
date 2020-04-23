/*
 *  errors.go
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

package utils

import (
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// RESTErrors contains unified REST error
type RESTErrors struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Errors  interface{} `json:"errors"`
}

// RESTError contains unified REST error
type RESTError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
	Error   string `json:"error"`
}

// RESTNoError contains unified REST error
type RESTNoError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

// BaseError creates unified error response
func BaseError(code int, message string, _type string, body interface{}) *echo.HTTPError {
	if body == nil {
		return echo.NewHTTPError(code, &RESTNoError{
			Code:    code,
			Message: message,
			Type:    _type,
		})
	}
	switch v := reflect.ValueOf(body); v.Kind() {
	case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct:
		return echo.NewHTTPError(code, &RESTErrors{
			Code:    code,
			Message: message,
			Type:    _type,
			Errors:  body,
		})
	case reflect.String:
		return echo.NewHTTPError(code, &RESTError{
			Code:    code,
			Message: message,
			Type:    _type,
			Error:   v.String(),
		})
	}
	return nil
}

// ProcessError handles errors based on their type
func ProcessError(err error) *echo.HTTPError {
	if os.Getenv("DEBUG_MODE") == "true" {
		log.Println(err)
	}
	switch e := err.(type) {
	case validator.ValidationErrors:
		return processValidationError(e)
	case *mysql.MySQLError:
		return BaseError(http.StatusInternalServerError, "Database error", "database_error", e.Number)
	}
	switch err {
	case gorm.ErrRecordNotFound:
		return BaseError(http.StatusNotFound, "Record Not Found", "database_error", err.Error())
	case bcrypt.ErrMismatchedHashAndPassword:
		return BaseError(http.StatusUnauthorized, "Invalid Credentials", "auth_error", err.Error())
	}
	return BaseError(http.StatusInternalServerError, "Unknown Error", "unknown_error", err.Error())
}

func processValidationError(errs validator.ValidationErrors) *echo.HTTPError {
	errsMap := make([]map[string]string, 0)
	for _, err := range errs {
		fieldErr := make(map[string]string)
		fieldErr["field"] = err.Field()
		fieldErr["tag"] = err.Tag()
		if err.Param() != "" {
			fieldErr["param"] = err.Param()
		}
		errsMap = append(errsMap, fieldErr)
	}
	return BaseError(http.StatusBadRequest, "Validation failed", "validation_failed", errsMap)
}
