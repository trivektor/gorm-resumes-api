package models

import (
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	ResumeId int32 `json:"resume_id"`
	Resume Resume
	Title string `json:"title"`
	CompanyName string `json:"company_name"`
	Description string `json:"description"`
	EmploymentType string `json:"employment_type"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	Industry string `json:"industry"`
	Location string `json:"location"`
}