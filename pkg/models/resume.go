package models

import (
	"gorm.io/gorm"
)

type Resume struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedBy string `json:"created_by"`
	Intro Intro
	Educations []Education
	Positions []Position
	References []Reference
	Skills []Skill
}
