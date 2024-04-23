package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/LCP/config"
	"github.com/toastsandwich/LCP/models"
	"github.com/toastsandwich/LCP/repository/query"
	"golang.org/x/crypto/bcrypt"
)

var sessionManager = config.SessionManager

func GetAllDoctors(ctx echo.Context) error {
	docs, err := query.ShowAllDoctors()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, docs)
}

func GetDoctorByID(ctx echo.Context) error {
	id := ctx.Param("id")
	doc, err := query.SearchDoctorByDocID(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, doc)
}

// func GetDoctorsByFirstNameLastName(ctx echo.Context) error {

// }

func AddDoctor(ctx echo.Context) error {
	var doc models.Doctor
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
	err = query.CreateDoctor(&doc)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, map[string]string{"message": "record created"})
}

func DeleteDoctor(ctx echo.Context) error {
	id := ctx.Param("id")
	err := query.RemoveDoctor(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "record deleted"})
}

func DoctorLogin(ctx echo.Context) error {
	docid := ctx.FormValue("docid")
	password := ctx.FormValue("password")
	doc := query.AttemptLogin_DOC(docid, password)
	if doc != nil {
		sessionManager.Put(ctx.Request().Context(), "firstname", doc.FirstName)
		sessionManager.Put(ctx.Request().Context(), "lastname", doc.LastName)
		// return ctx.JSON(http.StatusOK, doc)
		return ctx.Redirect(http.StatusFound, "http://localhost:8080/doctor/home?doc_id="+doc.DocID+"&first_name="+doc.FirstName+"&last_name="+doc.LastName)
	} else {
		return ctx.Redirect(http.StatusFound, "http://localhost:8080/doctor/login")
	}
}
