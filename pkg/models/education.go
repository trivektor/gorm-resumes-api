package models

import (
	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	ResumeId int32 `json:"resume_id"`
	Resume Resume
	Id int32 `json:"id"`
	School string `json:"school"`
	Degree string `json:"degree"`
	FieldOfStudy string `json:"field_of_study"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
}