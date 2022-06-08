package models

import (
	"gorm.io/gorm"
)

type Reference struct {
	gorm.Model
	ResumeID int32 `json:"resume_id"`
	Id int32 `json:"id"`
	FullName string `json:"full_name"`
	Company string `json:"company"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}