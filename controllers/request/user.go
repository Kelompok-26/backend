package request

import (
	"backend/helper"
	"backend/models"
)

type ReqUser struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Password      string `json:"password"`
	DateofBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	AccountNumber string `json:"account_number"`
}

func (user *ReqUser) MapToDomain() models.User {
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
