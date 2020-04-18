package database

import (
	"github.com/jinzhu/gorm"
	"github.com/metal3d/go-slugify"
)

type (
	// Tag represents a tag
	Tag struct {
		gorm.Model
		UserID uint
		Slug   string
		Name   string
		Links  []*Link `gorm:"many2many:link_tags;"`
	}

	// ResponseTag represents response version of Tag
	ResponseTag struct {
		Name  string `json:"name"`
		Slug  string `json:"slug"`
		Count uint   `json:"count"`
	}
)

// Count returns number of links containing Tag
func (tag *Tag) Count() uint {
	return uint(DB().Model(&tag).Association("Links").Count())
}

// GetResponseTag returns ResponseTag from Tag
func (tag *Tag) GetResponseTag() *ResponseTag {
	return &ResponseTag{
		Name:  tag.Name,
		Slug:  tag.Slug,
		Count: tag.Count(),
	}
}

// CreateTags inserts multiple tags into DB
func CreateTags(userID uint, str []string) ([]*Tag, error) {
	tags := make([]*Tag, 0)
	err := DB().Transaction(func(tx *gorm.DB) error {
		for _, t := range str {
			tag := &Tag{
				UserID: userID,
				Name:   t,
				Slug:   slugify.Marshal(t),
			}
			if err := tx.FirstOrCreate(&tag, *tag).Error; err != nil {
				return err
			}
			tags = append(tags, tag)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// GetTag retrieves Tag from DB
func GetTag(userID uint, slug string) (*Tag, error) {
	tag := new(Tag)
	err := DB().
		Where("user_id = ?", userID).
		Where("slug = ?", slug).First(tag).Error
	return tag, err
}

// GetTags retrieves all tags from DB
func GetTags(userID uint) ([]*Tag, error) {
	tags := make([]*Tag, 0)
	err := DB().
		Where("user_id = ?", userID).
		Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// GetLinks retrieves links containing Tag
func (tag *Tag) GetLinks() ([]*Link, error) {
	links := make([]*Link, 0)
	err := DB().Model(&tag).Preload("Tags").Related(&links, "Links").Error
	return links, err
}
