package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/LCP/models"
	"github.com/toastsandwich/LCP/repository/query"
	"golang.org/x/crypto/bcrypt"
)

func GetAllLabAssistants(ctx echo.Context) error {
	labAsst, err := query.ShowAllLabAssistants()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, labAsst)
}

func GetLabAssistantByID(ctx echo.Context) error {
	id := ctx.Param("id")
	asst, err := query.SearchLabAssistantByID(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, asst)
}

func AddLabAssistant(ctx echo.Context) error {
	var labAssistant models.LabAssistant
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
	err = query.CreateLabAssistant(&labAssistant)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, map[string]string{"message": "record created"})
}

func DeleteLabAssistant(ctx echo.Context) error {
	id := ctx.Param("id")
	err := query.RemoveLabAssistant(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "record removed"})
}

func LabAsstLogin(ctx echo.Context) error {
	labAsstID := ctx.FormValue("labAsstID")
	password := ctx.FormValue("password")
	labAsst := query.AttemptLogin_LABASST(labAsstID, password)
	fmt.Println(labAsstID)
	fmt.Println(labAsst)
	if labAsst != nil {
		return ctx.Redirect(http.StatusFound, "http://localhost:8080/lab-assistant/home?lab_asst_id="+labAsst.LabAsstID+"&first_name="+labAsst.FirstName+"&last_name="+labAsst.LastName)
	} else {
		return ctx.Redirect(http.StatusFound, "http://localhost:8080/lab-assistant/login")
	}
}
