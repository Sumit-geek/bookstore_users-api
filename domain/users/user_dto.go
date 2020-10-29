package users

import (
	"github.com/Sumit-geek/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:id`
	FirstName   string `json:"first_name""`
	LastName    string `json:"last_name""`
	Email       string `json:"email""`
	DateCreated string `json:"date_created""`
}

func (user *User) Validate(isPartial bool) *errors.RestErr {
	strings.TrimSpace(user.FirstName)
	strings.TrimSpace(user.LastName)
	
	if !isPartial {
		user.Email = strings.TrimSpace(strings.ToLower(user.Email))
		if user.Email == "" {
			return errors.NewBadRequestError("invalid email address")
		}
	}
	return nil
}