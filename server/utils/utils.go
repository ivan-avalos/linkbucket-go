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
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Token represents a JWT token
type Token struct {
	UserID uint
	jwt.StandardClaims
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
