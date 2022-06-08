package models

import (
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	ResumeID int32 `json:"resume_id"`
	Id int32 `json:"id"`
	Skill string `json:"skill"`
	Description string `json:"description"`
}