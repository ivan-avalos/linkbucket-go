/* 
 *  link-controller.go
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
	"strconv"

	"github.com/ivan-avalos/linkbucket/database"
	"github.com/ivan-avalos/linkbucket/utils"
	"github.com/labstack/echo"
)

func responseLinks(links []*database.Link) []*database.ResponseLink {
	rlinks := make([]*database.ResponseLink, 0)
	for _, l := range links {
		rlinks = append(rlinks, l.GetResponseLink())
	}
	return rlinks
}

func processTagsForLink(link *database.Link, str string) error {
	tags := utils.SeparateTagsFromString(str)
	if err := link.LinkTags(tags); err != nil {
		return utils.ProcessError(err)
	}
	return nil
}

// CreateLink creates link for user
func CreateLink(c echo.Context) error {
	userID := utils.GetJWTUserID(c)
	requestLink := new(database.RequestLink)
	if err := c.Bind(requestLink); err != nil {
		return utils.ProcessError(err)
	}
	if err := c.Validate(requestLink); err != nil {
		return utils.ProcessError(err)
	}

	link := requestLink.GetLink()
	link.UserID = userID
	if err := database.FieldIsUniqueForUser(userID, "links", "link", link.Link); err != nil {
		return err
	}
	if err := link.Create(); err != nil {
		return utils.ProcessError(err)
	}
	if err := processTagsForLink(link, requestLink.Tags); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, link.GetResponseLink())
}

// GetLink returns link with ID
func GetLink(c echo.Context) error {
	userID := utils.GetJWTUserID(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ProcessError(err)
	}
	link, err := database.GetLink(uint(id), userID)
	if err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, link.GetResponseLink())
}

// GetLinks returns all user links
func GetLinks(c echo.Context) error {
	userID := utils.GetJWTUserID(c)
	links, err := database.GetLinks(userID)
	if err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, responseLinks(links))
}

// GetLinksForTag returns all links containing Tag
func GetLinksForTag(c echo.Context) error {
	userID := utils.GetJWTUserID(c)
	slug := c.Param("slug")
	tag, err := database.GetTag(userID, slug)
	if err != nil {
		return utils.ProcessError(err)
	}
	links, err := tag.GetLinks()
	if err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, responseLinks(links))
}

// GetLinksForSearch returns links for search query
func GetLinksForSearch(c echo.Context) error {
	type search struct {
		Query string `json:"query" validate:"required"`
	}
	s := new(search)
	if err := c.Bind(s); err != nil {
		return utils.ProcessError(err)
	}
	if err := c.Validate(s); err != nil {
		return utils.ProcessError(err)
	}
	userID := utils.GetJWTUserID(c)
	links, err := database.GetLinksForSearch(userID, s.Query)
	if err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, responseLinks(links))
}

// UpdateLink modifies link in DB
func UpdateLink(c echo.Context) error {
	userID := utils.GetJWTUserID(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ProcessError(err)
	}
	updateLink := new(database.UpdateLink)
	if err := c.Bind(updateLink); err != nil {
		return utils.ProcessError(err)
	}
	if err := c.Validate(updateLink); err != nil {
		return utils.ProcessError(err)
	}
	link, err := database.GetLink(uint(id), userID)
	if err != nil {
		return utils.ProcessError(err)
	}
	link.UserID = userID
	if updateLink.Title != "" {
		link.Title = updateLink.Title
	}
	if updateLink.Link != "" {
		if err := database.FieldIsSameOrUniqueForUser(userID, "links", uint(id), "link", updateLink.Link); err != nil {
			return err
		}
		link.Link = updateLink.Link
	}
	if updateLink.Tags == "-" {
		if err := link.UnlinkTags(); err != nil {
			return err
		}
	}
	if updateLink.Tags != "" && updateLink.Tags != "-" {
		if err := processTagsForLink(link, updateLink.Tags); err != nil {
			return err
		}
	}
	if err := link.Update(); err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, link.GetResponseLink())
}

// DeleteLink removes Link from DB
func DeleteLink(c echo.Context) error {
	userID := utils.GetJWTUserID(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ProcessError(err)
	}
	link, err := database.GetLink(uint(id), userID)
	if err != nil {
		return utils.ProcessError(err)
	}
	if err := link.Delete(); err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, link.GetResponseLink())
}
