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
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/badoux/goscraper"
	"github.com/ivan-avalos/gorm-paginator/pagination"
	"github.com/ivan-avalos/linkbucket-go/server/database"
	"github.com/ivan-avalos/linkbucket-go/server/jobs"
	"github.com/ivan-avalos/linkbucket-go/server/utils"
	"github.com/labstack/echo/v4"
)

func responseLinks(links []*database.Link) utils.Response {
	rlinks := make([]*database.ResponseLink, 0)
	for _, l := range links {
		rlinks = append(rlinks, l.GetResponseLink())
	}
	return utils.BaseResponse(http.StatusOK, rlinks)
}

func responseLinksPag(links []*database.Link, pag *pagination.Paginator) utils.Paginate {
	rlinks := make([]interface{}, 0)
	for _, l := range links {
		rlinks = append(rlinks, l.GetResponseLink())
	}
	return utils.PaginateResponse(http.StatusOK, pag, rlinks)
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
	// Get <title> from URL if found
	if s, err := goscraper.Scrape(link.Title, 5); err == nil {
		link.Title = s.Preview.Title
	}
	if err := link.Create(); err != nil {
		return utils.ProcessError(err)
	}
	if err := processTagsForLink(link, requestLink.Tags); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, utils.BaseResponse(http.StatusOK, link.GetResponseLink()))
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
	return c.JSON(http.StatusOK, utils.BaseResponse(http.StatusOK, link.GetResponseLink()))
}

// GetLinks returns all user links
func GetLinks(c echo.Context) error {
	userID := utils.GetJWTUserID(c)
	p := new(database.Paginate)
	if err := c.Bind(p); err != nil {
		return utils.ProcessError(err)
	}
	if p.Page != 0 && p.Limit != 0 {
		links, pag, err := database.GetLinks(userID, int(p.Page), int(p.Limit))
		if err != nil {
			return utils.ProcessError(err)
		}
		return c.JSON(http.StatusOK, responseLinksPag(links, pag))
	}

	links, _, err := database.GetLinks(userID, 0, -1)
	if err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, responseLinks(links))
}

// GetLinksForTag returns links containing tag
func GetLinksForTag(c echo.Context) error {
	userID := utils.GetJWTUserID(c)
	slug := c.Param("slug")
	// Get Tag
	tag, err := database.GetTag(userID, slug)
	if err != nil {
		return utils.ProcessError(err)
	}
	p := new(database.Paginate)
	if err := c.Bind(p); err != nil {
		return utils.ProcessError(err)
	}
	if p.Page != 0 && p.Limit != 0 {
		links, pag, err := tag.GetLinks(int(p.Page), int(p.Limit))
		if err != nil {
			return utils.ProcessError(err)
		}
		return c.JSON(http.StatusOK, responseLinksPag(links, pag))
	}

	links, _, err := tag.GetLinks(0, -1)
	if err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, responseLinks(links))
}

// GetLinksForSearch returns links for search query
func GetLinksForSearch(c echo.Context) error {
	userID := utils.GetJWTUserID(c)
	s := new(database.Search)
	if err := c.Bind(s); err != nil {
		return utils.ProcessError(err)
	}
	if err := c.Validate(s); err != nil {
		return utils.ProcessError(err)
	}
	if s.Page != 0 && s.Limit != 0 {
		links, pag, err := database.GetLinksForSearch(userID, s.Query, int(s.Page), int(s.Limit))
		if err != nil {
			return utils.ProcessError(err)
		}
		return c.JSON(http.StatusOK, responseLinksPag(links, pag))
	}

	links, _, err := database.GetLinksForSearch(userID, s.Query, 0, -1)
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
	return c.JSON(http.StatusOK, utils.BaseResponse(http.StatusOK, link.GetResponseLink()))
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
	return c.JSON(http.StatusOK, utils.BaseResponse(http.StatusOK, link.GetResponseLink()))
}

// ImportOld imports JSON file from old Linkbucket
func ImportOld(c echo.Context) error {
	userID := utils.GetJWTUserID(c)

	// Upload
	file, err := c.FormFile("links")
	if err != nil {
		return utils.ProcessError(err)
	}
	src, err := file.Open()
	if err != nil {
		return utils.ProcessError(err)
	}
	defer src.Close()

	var links []database.OldImportLink
	if err := json.NewDecoder(src).Decode(&links); err != nil {
		return utils.ProcessError(err)
	}
	job := database.Job{
		UserID: userID,
		Name:   "AsyncOldImport",
		FnName: "jobs.AsyncOldImport",
		Params: []interface{}{links},
	}
	if err := job.Create(); err != nil {
		return utils.ProcessError(err)
	}
	if err := jobs.Enqueue(job); err != nil {
		return utils.ProcessError(err)
	}

	return c.JSON(http.StatusOK, utils.BaseResponse(http.StatusOK, job.GetResponseJob()))
}

// ImportNew imports JSON file from Linkbucket Go
func ImportNew(c echo.Context) error {
	return nil
}
