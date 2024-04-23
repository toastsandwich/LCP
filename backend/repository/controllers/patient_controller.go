package controllers

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/LCP/models"
	"github.com/toastsandwich/LCP/repository/query"
)

func GetAllPatients(ctx echo.Context) error {
	paitents, err := query.ShowAllPatients()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, paitents)
}

func GetPatientByID(ctx echo.Context) error {
	id := ctx.Param("id")
	patient, err := query.SelectPatitentByID(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, patient)
}

type Prediction struct {
	Selector []string `json:"selector"`
}

func CreatePatient(ctx echo.Context) error {

	convFloat32 := func(val1 string, val2 *float32) error {
		val, err := strconv.ParseFloat(val1, 32)
		if err != nil {
			return err
		}
		*val2 = float32(val)
		return nil
	}
	convUInt := func(val1 string, val2 *uint) error {
		val, err := strconv.ParseUint(val1, 10, 64)
		if err != nil {
			return err
		}
		*val2 = uint(val)
		return nil
	}

	var patient models.Patient
	var patientreq models.PatientRequest
	err := ctx.Bind(&patientreq)
	if err != nil {
		return err
	}

	patient.FirstName = patientreq.FirstName
	patient.LastName = patientreq.LastName
	patient.PatientID = patient.FirstName + patient.LastName
	patient.Gender = patientreq.Gender
	patient.EMail = patientreq.EMail
	patient.PhoneNumber = patientreq.PhoneNumber
	patient.Selector = uint(rand.Intn(2) + 1)
	convUInt(patientreq.Age, &patient.Age)
	convFloat32(patientreq.TB, &patient.TB)
	convFloat32(patientreq.DB, &patient.DB)
	convFloat32(patientreq.Alkphos, &patient.Alkphos)
	convFloat32(patientreq.SGOT, &patient.SGOT)
	convFloat32(patientreq.SGPT, &patient.SGPT)
	convFloat32(patientreq.ALB, &patient.ALB)
	convFloat32(patientreq.AG_RATIO, &patient.AG_RATIO)

	err = query.CreatePatient(&patient)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "record created"})
}

func DeletePatient(ctx echo.Context) error {
	id := ctx.Param("id")
	err := query.RemovePatient(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "record removed"})
}
