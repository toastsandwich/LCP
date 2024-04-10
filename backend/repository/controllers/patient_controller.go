package controllers

import (
	"net/http"

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

func CreatePatient(ctx echo.Context) error {
	var patient models.Patient
	err := ctx.Bind(&patient)
	if err != nil {
		return err
	}
	err = query.CreatePatient(patient)
	if err != nil {
		return err
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
