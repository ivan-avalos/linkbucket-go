/* 
 *  link.go
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

package database

import (
	"github.com/jinzhu/gorm"
)

type (
	// Link represents a link
	Link struct {
		gorm.Model
		UserID uint
		Title  string `gorm:"type:mediumtext;not null"`
		Link   string `gorm:"type:mediumtext;not null"`
		Tags   []*Tag `gorm:"many2many:link_tags;"`
	}

	// RequestLink represents request version of Link
	RequestLink struct {
		Title string `json:"title" validate:"lt=255"`
		Link  string `json:"link" validate:"required,url,lt=255"`
		Tags  string `json:"tags" validate:"tags"`
	}

	// UpdateLink represents request version of Link
	UpdateLink struct {
		Title string `json:"title" validate:"lt=255"`
		Link  string `json:"link" validate:"omitempty,url,lt=255"`
		Tags  string `json:"tags" validate:"omitempty,tags"`
	}

	// ResponseLink represents response version of Link
	ResponseLink struct {
		ID    uint          `json:"id"`
		Title string        `json:"title"`
		Link  string        `json:"link"`
		Tags  []ResponseTag `json:"tags"`
	}
)

// GetLink returns Link from RequestLink
func (rl *RequestLink) GetLink() *Link {
	return &Link{
		Title: rl.Title,
		Link:  rl.Link,
	}
}

// GetLink returns Link from UpdateLink
func (ul *UpdateLink) GetLink() *Link {
	return &Link{
		Title: ul.Title,
		Link:  ul.Link,
	}
}

// GetResponseLink returns ResponseLink from Link
func (link *Link) GetResponseLink() *ResponseLink {
	responseTags := make([]ResponseTag, 0)
	for _, t := range link.Tags {
		if t.Count() > 0 {
			responseTags = append(responseTags, *t.GetResponseTag())
		}
	}
	return &ResponseLink{
		ID:    link.ID,
		Title: link.Title,
		Link:  link.Link,
		Tags:  responseTags,
	}
}

// Create inserts a new link into DB
func (link *Link) Create() error {
	return DB().Create(link).Error
}

// LinkTags inserts multiple Tags to Link
func (link *Link) LinkTags(str []string) error {
	tags, err := CreateTags(link.UserID, str)
	if err != nil {
		return err
	}
	intags := make([]interface{}, len(tags))
	for i, t := range tags {
		intags[i] = t
	}
	err = DB().Model(&link).Association("Tags").Replace(intags...).Error
	if err != nil {
		return err
	}
	link.Tags = tags
	return nil
}

// UnlinkTags clears all Tags from Link
func (link *Link) UnlinkTags() error {
	return DB().Model(&link).Association("Tags").Clear().Error
}

// GetLink retrieves a link from DB
func GetLink(id, userID uint) (*Link, error) {
	link := &Link{}
	err := DB().Table("links").
		Where("id = ?", id).
		Where("user_id = ?", userID).
		Preload("Tags").
		First(link).Error
	return link, err
}

// GetLinks retrieves all user links from DB
func GetLinks(userID uint) ([]*Link, error) {
	links := make([]*Link, 0)
	err := DB().Where("user_id = ?", userID).Preload("Tags").Find(&links).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}

// GetLinksForSearch retrieves all user links containing query
func GetLinksForSearch(userID uint, q string) ([]*Link, error) {
	links := make([]*Link, 0)
	err := DB().Where("user_id = ?", userID).
		Where("title LIKE ? OR link LIKE ?", "%"+q+"%", "%"+q+"%").
		Preload("Tags").Find(&links).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}

// Update modifies link in DB
func (link *Link) Update() error {
	return DB().Save(link).Error
}

// Delete removes link from DB
func (link *Link) Delete() error {
	return DB().Delete(link).Error
}
