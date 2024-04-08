package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	PatientID string  `json: "patient_id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Gender    string  `json:"gender"`
	Age       uint    `json:"age"`
	TB        float32 `json:"tb"`
	DB        float32 `json:"db"`
	Alkphos   float32 `json:"alkphos"`
	SGPT      float32 `json:"sgpt"`
	SGOT      float32 `json:"sgot"`
	ALB       float32 `json:"alb"`
	AG_RATIO  float32 `json:"a/g_ratio"`
	Selector  uint    `json:"selector"`
}

func init() {
	db.AutoMigrate(&Patient{})
}

// SQL

func showAllPatients() ([]*Patient, error) {
	var patients []*Patient
	err := db.Model(&Patient{}).Select("id", "patient_id", "first_name", "last_name", "gender", "age", "tb", "db", "alkphos", "sgpt", "sgot", "alb", "ag_ratio", "selector").Select(&patients).Error
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func selectPatitentByID(patient_id string) (*Patient, error) {
	var patient *Patient
	err := db.Where("patient_id = ?", patient_id).Find(&patient).Error
	if err != nil {
		return nil, err
	}
	return patient, nil
}

func createPatient(patient Patient) error {
	return db.Create(patient).Error
}

func removePatient(patient_id string) error {
	var patient Patient
	return db.Where("patient_id = ?", patient_id).Delete(&patient).Error
}

// ECHO

func GetAllPatients(ctx echo.Context) error {
	paitents, err := showAllPatients()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, paitents)
}

func GetPatientByID(ctx echo.Context) error {
	id := ctx.Param("id")
	patient, err := selectPatitentByID(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, patient)
}

func CreatePatient(ctx echo.Context) error {
	var patient Patient
	err := ctx.Bind(&patient)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "record created"})
}

func DeletePatient(ctx echo.Context) error {
	id := ctx.Param("id")
	err := removeDoctor(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "record removed"})
}
