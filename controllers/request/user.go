package request

import (
	"backend/helper"
	"backend/models"
)

type ReqNewUser struct {
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	PhoneNumber   string `json:"phone_number" validate:"required,min=10"`
	Password      string `json:"password" validate:"required,min=8"`
	DateofBirth   string `json:"date_of_birth" validate:"required"`
	Gender        string `json:"gender" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
}
type ReqUser struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Password      string `json:"password"`
	DateofBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	AccountNumber string `json:"account_number"`
}

func (user *ReqUser) MapToUser() models.User {
	return models.User{
		Name:          user.Name,
		PhoneNumber:   user.PhoneNumber,
		Email:         user.Email,
		Password:      helper.CreateHash(user.Password),
		DateofBirth:   helper.ConvertStringToDate(user.DateofBirth),
		Gender:        user.Gender,
		AccountNumber: user.AccountNumber,
	}
}

func (user *ReqNewUser) MapToNewUser() models.User {
	return models.User{
		Name:          user.Name,
		PhoneNumber:   user.PhoneNumber,
		Email:         user.Email,
		Password:      helper.CreateHash(user.Password),
		DateofBirth:   helper.ConvertStringToDate(user.DateofBirth),
		Gender:        user.Gender,
		AccountNumber: user.AccountNumber,
	}
}
