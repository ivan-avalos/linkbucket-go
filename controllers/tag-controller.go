package controllers

import (
	"net/http"

	"github.com/ivan-avalos/linkbucket/database"
	"github.com/ivan-avalos/linkbucket/utils"
	"github.com/labstack/echo"
)

func responseTags(links []*database.Tag) []*database.ResponseTag {
	rtags := make([]*database.ResponseTag, 0)
	for _, t := range links {
		rtags = append(rtags, t.GetResponseTag())
	}
	return rtags
}

// GetTags returns all Tags for User
func GetTags(c echo.Context) error {
	userID := utils.GetJWTUserID(c)
	tags, err := database.GetTags(userID)
	if err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, responseTags(tags))
}
