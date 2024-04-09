package models

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/LCP/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db = config.DB

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

// MYSQL

func init() {
	db.AutoMigrate(&Doctor{})
}

func showAllDoctors() ([]*Doctor, error) {
	var doctors []*Doctor
	err := db.Model(&Doctor{}).Select("DocID", "FirstName", "LastName", "Gender", "age", "Email", "PhoneNumber").Find(&doctors).Error
	if err != nil {
		return nil, err
	}
	return doctors, nil
}

func searchDoctorByDocID(docID string) (Doctor, error) {
	var doctor Doctor
	err := db.Where("doc_id = ?", docID).Find(&doctor).Error
	if err != nil {
		return Doctor{}, err
	}
	return doctor, nil
}

func createDoctor(newDoc *Doctor) error {
	return db.Create(newDoc).Error
}

func removeDoctor(docID string) error {
	var doctor Doctor
	fmt.Println(doctor)
	return db.Where("doc_id = ?", docID).Delete(&doctor).Error
}

// ECHO

func GetAllDoctors(ctx echo.Context) error {
	docs, err := showAllDoctors()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, docs)
}

func GetDoctorByID(ctx echo.Context) error {
	id := ctx.Param("id")
	doc, err := searchDoctorByDocID(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, doc)
}

// func GetDoctorsByFirstNameLastName(ctx echo.Context) error {

// }

func AddDoctor(ctx echo.Context) error {
	var doc Doctor
	err := ctx.Bind(&doc)
	if err != nil {
		return err
	}
	pass := doc.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	doc.Password = string(hash)
	err = createDoctor(&doc)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, map[string]string{"message": "record created"})
}

func DeleteDoctor(ctx echo.Context) error {
	id := ctx.Param("id")
	err := removeDoctor(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "record deleted"})
}
