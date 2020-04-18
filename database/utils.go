package database

import (
	"net/http"

	"github.com/ivan-avalos/linkbucket/utils"
	"github.com/jinzhu/gorm"
)

// FieldIsUniqueForUser returns err if field is not unique for user
func FieldIsUniqueForUser(userID uint, table string, field string, value string) error {
	count := 0
	err := DB().Table(table).
		Where("user_id = ?", userID).
		Where(field+" = ?", value).
		Count(&count).Error
	if err != nil {
		return utils.ProcessError(err)
	}
	if count != 0 {
		return utils.BaseError(
			http.StatusBadRequest,
			"Validation failed",
			"validation_failed",
			[]map[string]string{
				{
					"field": field,
					"tag":   "unique",
				},
			},
		)
	}
	return nil
}

// FieldIsSameOrUnique returns err if new value is not different and not unique
func FieldIsSameOrUnique(table string, id uint, field string, new string) error {
	obj := &gorm.Model{}
	err := DB().Table(table).
		Where(field+" = ?", new).First(obj).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return utils.ProcessError(err)
	}
	if obj.ID != id {
		return utils.BaseError(
			http.StatusBadRequest,
			"Validation failed",
			"validation_failed",
			[]map[string]string{
				{
					"field": field,
					"tag":   "unique",
				},
			},
		)
	}
	return nil
}

// FieldIsSameOrUniqueForUser returns err if new value is not different and not unique for user
func FieldIsSameOrUniqueForUser(userID uint, table string, id uint, field string, new string) error {
	obj := &gorm.Model{}
	err := DB().Table(table).
		Where("user_id = ?", userID).
		Where(field+" = ?", new).First(obj).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return utils.ProcessError(err)
	}
	if obj.ID != id {
		return utils.BaseError(
			http.StatusBadRequest,
			"Validation failed",
			"validation_failed",
			[]map[string]string{
				{
					"field": field,
					"tag":   "unique",
				},
			},
		)
	}
	return nil
}
