package query

import (
	"github.com/toastsandwich/LCP/config"
	"github.com/toastsandwich/LCP/models"
	"golang.org/x/crypto/bcrypt"
)

var db = config.DB

func init() {
	db.AutoMigrate(&models.Doctor{})
}

func ShowAllDoctors() ([]*models.Doctor, error) {
	var doctors []*models.Doctor
	err := db.Model(&models.Doctor{}).Select("DocID", "FirstName", "LastName", "Gender", "age", "Email", "PhoneNumber").Find(&doctors).Error
	if err != nil {
		return nil, err
	}
	return doctors, nil
}

func SearchDoctorByDocID(docID string) (models.Doctor, error) {
	var doctor models.Doctor
	err := db.Where("doc_id = ?", docID).Find(&doctor).Error
	if err != nil {
		return models.Doctor{}, err
	}
	return doctor, nil
}

func AttemptLogin_DOC(docID, password string) *models.Doctor {
	var possibles []models.Doctor
	db.Where("doc_id = ?", docID).Find(&possibles)
	for _, doc := range possibles {
		err := bcrypt.CompareHashAndPassword([]byte(doc.Password), []byte(password))
		if err == nil {
			return &doc
		}
	}
	return nil
}

func CreateDoctor(newDoc *models.Doctor) error {
	return db.Create(newDoc).Error
}

func RemoveDoctor(docID string) error {
	var doctor models.Doctor
	return db.Where("doc_id = ?", docID).Delete(&doctor).Error
}
