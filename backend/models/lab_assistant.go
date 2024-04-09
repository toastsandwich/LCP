package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

// MYSQL
func init() {
	db.AutoMigrate(&LabAssistant{})
}
func showAllLabAssistants() ([]*LabAssistant, error) {
	var labAssistants []*LabAssistant
	err := db.Model(&LabAssistant{}).Select("LabAsstID", "FirstName", "LastName").Select(&labAssistants).Error
	if err != nil {
		return nil, err
	}
	return labAssistants, nil
}

func searchLabAssistantByID(lab_asst_id string) (*LabAssistant, error) {
	var labAssistant *LabAssistant
	err := db.Where("lab_asst_id = ?", lab_asst_id).Find(&labAssistant).Error
	if err != nil {
		return nil, err
	}
	return labAssistant, nil
}

func createLabAssistant(newEntry *LabAssistant) error {
	return db.Create(newEntry).Error
}

func removeLabAssistant(lab_asst_id string) error {
	var labAssistant LabAssistant
	return db.Where("lab_asst_id = ?", lab_asst_id).Delete(&labAssistant).Error
}

// ECHO

func GetAllLabAssistants(ctx echo.Context) error {
	labAsst, err := showAllLabAssistants()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, labAsst)
}

func GetLabAssistantByID(ctx echo.Context) error {
	id := ctx.Param("id")
	asst, err := searchLabAssistantByID(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, asst)
}

func AddLabAssistant(ctx echo.Context) error {
	var labAssistant LabAssistant
	err := ctx.Bind(&labAssistant)
	if err != nil {
		return err
	}
	pass := labAssistant.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	labAssistant.Password = string(hash)
	err = createLabAssistant(&labAssistant)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, map[string]string{"message": "record created"})
}

func DeleteLabAssistant(ctx echo.Context) error {
	id := ctx.Param("id")
	err := removeLabAssistant(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "record removed"})
}
