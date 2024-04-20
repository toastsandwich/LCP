package query

import "github.com/toastsandwich/LCP/models"

func init() {
	db.AutoMigrate(&models.Patient{})
}

func ShowAllPatients() ([]*models.Patient, error) {
	var patients []*models.Patient
	err := db.Model(&models.Patient{}).Select("patient_id", "first_name", "last_name", "gender", "phone_number", "e_mail", "age", "tb", "db", "alkphos", "sgpt", "sgot", "alb", "ag_ratio", "selector").Find(&patients).Error
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func SelectPatitentByID(patient_id string) (*models.Patient, error) {
	var patient *models.Patient
	err := db.Where("patient_id = ?", patient_id).Find(&patient).Error
	if err != nil {
		return nil, err
	}
	return patient, nil
}

func CreatePatient(patient models.Patient) error {
	return db.Create(patient).Error
}

func RemovePatient(patient_id string) error {
	var patient models.Patient
	return db.Where("patient_id = ?", patient_id).Delete(&patient).Error
}
