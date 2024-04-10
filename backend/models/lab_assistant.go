package models

import (
	"gorm.io/gorm"
)

type LabAssistant struct {
	gorm.Model
	LabAsstID   string `json:"lab_asst_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         uint   `json:"age"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
