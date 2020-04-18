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
