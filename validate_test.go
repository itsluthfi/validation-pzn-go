package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	user := ""

	if err := validate.Var(user, "required"); err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationTwoVariables(t *testing.T) {
	validate := validator.New()
	password := "rahasia"
	confirmPassword := "bukanrahasia"

	if err := validate.VarWithValue(password, confirmPassword, "eqfield"); err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()
	user := "itsluthfi"

	if err := validate.Var(user, "required,numeric"); err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	user := "01234567890"

	if err := validate.Var(user, "required,numeric,min=5,max=10"); err != nil {
		fmt.Println(err.Error())
	}
}

type LoginRequest struct {
	Username string `validate:"required,email"`
	Password string `validate:"required,min=5"`
}

func TestValidationStruct(t *testing.T) {
	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "itsluthfi",
		Password: "rhsa",
	}

	if err := validate.Struct(loginRequest); err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "itsluthfi",
		Password: "rhsa",
	}

	if err := validate.Struct(loginRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}
