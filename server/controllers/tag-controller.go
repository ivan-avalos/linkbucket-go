/*
 *  tag-controller.go
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

package controllers

import (
	"net/http"

	"github.com/ivan-avalos/linkbucket-go/server/database"
	"github.com/ivan-avalos/linkbucket-go/server/utils"
	"github.com/labstack/echo/v4"
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
	return c.JSON(http.StatusOK, utils.BaseResponse(http.StatusOK, responseTags(tags)))
}
