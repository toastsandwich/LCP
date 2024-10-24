package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	PatientID   string  `json:"patient_id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Gender      string  `json:"gender"`
	PhoneNumber string  `json:"phone_number"`
	EMail       string  `json:"email"`
	Age         uint    `json:"age"`
	TB          float32 `json:"tb"`
	DB          float32 `json:"db"`
	Alkphos     float32 `json:"alkphos"`
	SGPT        float32 `json:"sgpt"`
	SGOT        float32 `json:"sgot"`
	ALB         float32 `json:"alb"`
	AG_RATIO    float32 `json:"a/g_ratio"`
	Selector    uint    `json:"selector"`
}

type PatientRequest struct {
	PatientID   string `json:"patient_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	EMail       string `json:"email"`
	Age         string `json:"age"`
	TB          string `json:"tb"`
	DB          string `json:"db"`
	Alkphos     string `json:"alkphos"`
	SGPT        string `json:"sgpt"`
	SGOT        string `json:"sgot"`
	ALB         string `json:"alb"`
	AG_RATIO    string `json:"a/g_ratio"`
}
