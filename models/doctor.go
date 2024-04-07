package models

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/LCP/config"
	"gorm.io/gorm"
)

var db = config.DB

type Doctor struct {
	gorm.Model
	DocID     string `json:"doc_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

// MYSQL

func init() {
	db.AutoMigrate(&Doctor{})
}

func showAllDoctors() ([]*Doctor, error) {
	var doctors []*Doctor
	err := db.Model(&Doctor{}).Select("DocID", "FirstName", "LastName").Find(&doctors).Error
	if err != nil {
		return nil, err
	}
	return doctors, nil
}

// func SearchDoctorByName(firstname, lastname string) ([]*Doctor, error) {
// 	var doctors []*Doctor
// 	err := db.Where("FirstName = ? AND LastName = ?", firstname, lastname).Find(&doctors).Error
// 	if err != nil {
// 		return nil, err
// 	}connector
// 	return doctors, nil
// }

func searchDoctorByDocID(docID string) (Doctor, error) {
	var doctor Doctor
	err := db.Where("DocID = ?", docID).Find(&doctor).Error
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
	return db.Where("DocID = ?", docID).Delete(&doctor).Error
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
	fmt.Println("done")
	err := ctx.Bind(&doc)
	if err != nil {
		return err
	}
	fmt.Println("done")
	err = createDoctor(&doc)
	if err != nil {
		return err
	}
	fmt.Println("done")
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
