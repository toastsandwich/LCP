package models

import (
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	DocID       string `json:"doc_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         uint   `json:"age"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
