/*
 *  utils.go
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
	"math"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ivan-avalos/gorm-paginator/pagination"
	"github.com/labstack/echo/v4"
)

// Token represents a JWT token
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// Response represents standard response
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Paginate represents standard response for pagination
type Paginate struct {
	Code        int         `json:"code"`
	Total       int         `json:"total"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	FirstPage   int         `json:"first_page"`
	LastPage    int         `json:"last_page"`
	NextPage    int         `json:"next_page"`
	PrevPage    int         `json:"prev_page"`
	Data        interface{} `json:"data"`
}

// GetPagesNumber returns number of pages based on count and limit
func GetPagesNumber(count uint, limit uint) uint {
	return uint(math.Abs(float64(count / limit)))
}

// BaseResponse returns Response with data
func BaseResponse(code int, obj interface{}) Response {
	return Response{
		Code: code,
		Data: obj,
	}
}

// PaginateResponse returns PaginateResponse with data
func PaginateResponse(code int, pag *pagination.Paginator, obj interface{}) Paginate {
	return Paginate{
		Code:        code,
		Total:       pag.TotalPage,
		PerPage:     pag.Limit,
		CurrentPage: pag.Page,
		FirstPage:   pag.FirstPage,
		LastPage:    pag.LastPage,
		NextPage:    pag.NextPage,
		PrevPage:    pag.PrevPage,
		Data:        obj,
	}
}

// SeparateTagsFromString splits tags into []string
func SeparateTagsFromString(strtags string) []string {
	tags := make([]string, 0)
	for _, t := range strings.Split(strtags, ",") {
		tags = append(tags, strings.TrimSpace(t))
	}
	return RemoveDuplicatesFromSlice(tags)
}

// RemoveDuplicatesFromSlice is self-descriptive
func RemoveDuplicatesFromSlice(s []string) []string {
	m := make(map[string]bool)
	for _, item := range s {
		if _, ok := m[item]; !ok {
			m[item] = true
		}
	}

	var result []string
	for item := range m {
		result = append(result, item)
	}
	return result
}

// GetJWTUserID returns authenticated user ID from Context
func GetJWTUserID(c echo.Context) uint {
	return c.Get("user").(*jwt.Token).Claims.(*Token).UserID
}
