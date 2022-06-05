package models

import (
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	ResumeId int32 `json:"resume_id"`
	Resume Resume
	Id int32 `json:"id"`
	Skill string `json:"skill"`
	Description string `json:"description"`
}