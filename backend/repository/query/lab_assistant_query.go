package query

import (
	"github.com/toastsandwich/LCP/models"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	db.AutoMigrate(&models.LabAssistant{})
}
func ShowAllLabAssistants() ([]*models.LabAssistant, error) {
	var labAssistants []*models.LabAssistant
	err := db.Model(&models.LabAssistant{}).Select("LabAsstID", "FirstName", "LastName", "Gender", "age", "Email", "PhoneNumber").Find(&labAssistants).Error
	if err != nil {
		return nil, err
	}
	return labAssistants, nil
}

func SearchLabAssistantByID(lab_asst_id string) (*models.LabAssistant, error) {
	var labAssistant *models.LabAssistant
	err := db.Where("lab_asst_id = ?", lab_asst_id).Find(&labAssistant).Error
	if err != nil {
		return nil, err
	}
	return labAssistant, nil
}

func CreateLabAssistant(newEntry *models.LabAssistant) error {
	return db.Create(newEntry).Error
}

func RemoveLabAssistant(lab_asst_id string) error {
	var labAssistant models.LabAssistant
	return db.Where("lab_asst_id = ?", lab_asst_id).Delete(&labAssistant).Error
}

func AttemptLogin_LABASST(labAsstID, password string) *models.LabAssistant {
	var possibles []models.LabAssistant
	db.Where("lab_asst_id = ?", labAsstID).Find(&possibles)
	for _, doc := range possibles {
		err := bcrypt.CompareHashAndPassword([]byte(doc.Password), []byte(password))
		if err == nil {
			return &doc
		}
	}
	return nil
}
