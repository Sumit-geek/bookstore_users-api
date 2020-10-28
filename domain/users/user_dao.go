package users

import (
	"fmt"
	"github.com/Sumit-geek/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFound(fmt.Sprintf("User %d not found", user.Id))
	}
	
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", user.Id))
	}
	userDB[user.Id] = user
	return nil
}