package models

import (
	"gorm.io/gorm"
)

type Resume struct {
	gorm.Model
	Id int32 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedBy string `json:"created_by"`
}