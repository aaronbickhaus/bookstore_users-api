package users

import (
	"github.com/aaronbickhaus/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
	Status string `json:"status"`
	Password string `json:"password"`
}

type Users [] User

const (
	StatusActive = "active"
)
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
    user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))

	user.Status = strings.TrimSpace(strings.ToLower(user.Status))

	if user.Email == "" {
		return  errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return  errors.NewBadRequestError("invalid password")
	}

	return nil
}