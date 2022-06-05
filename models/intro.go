package models

import (
	"gorm.io/gorm"
)

type Intro struct {
	gorm.Model
	ResumeId int32 `json:"resume_id"`
	Resume Resume
	Id int32 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Headline string `json:"headline"`
	Industry string `json:"industry"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Summary string `json:"summary"`
	NamePronunciation string `json:"name_pronunciation"`
}