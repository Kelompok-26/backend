package response

import (
	"backend/helper"
	"backend/models"
)

type ResUser struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Password      string `json:"password"`
	DateofBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	Point         int    `json:"point"`
	AccountNumber string `json:"account_number"`
}

func MapToUser(user models.User) ResUser {
	return ResUser{
		Id:            user.Id,
		Name:          user.Name,
		PhoneNumber:   user.PhoneNumber,
		Email:         user.Email,
		Password:      user.Password,
		DateofBirth:   helper.ConvertDateToString(user.DateofBirth),
		Gender:        user.Gender,
		Point:         user.Point,
		AccountNumber: user.AccountNumber,
	}
}
