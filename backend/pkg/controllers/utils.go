package controllers

import (
	"EduConnect/internal/values"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(userHashedPassword values.Password, incomingPassword string) error {
	err := bcrypt.CompareHashAndPassword(userHashedPassword.GetPassword(), []byte(incomingPassword))
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}
	return nil
}

func (a *AuthController) decodeRequest(ctx echo.Context, i interface{}) error {
	if err := ctx.Bind(i); err != nil {
		return fmt.Errorf("invalid request")
	}

	if err := a.validate.Struct(i); err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}
