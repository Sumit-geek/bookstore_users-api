package users

import (
	"fmt"
	"github.com/Sumit-geek/bookstore_users-api/datasources/mysql/users_db"
	"github.com/Sumit-geek/bookstore_users-api/utils"
	"github.com/Sumit-geek/bookstore_users-api/utils/errors"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) values(?, ?, ?, ?)"
	queryReadUser = "SELECT id, first_name, last_name, email, date_created FROM USERS WHERE id = ?"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryReadUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("error while trying to get user %d:, %s", user.Id,
			err.Error()))
	}
	
	return nil
}

func (user *User) Save() *errors.RestErr {
	user.DateCreated = utils.GetNowString()
	
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	
	if saveErr != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exist", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error while trying to save user: %s", err.Error()))
	}
	
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error while trying to save user: %s", err.Error()))
	}
	
	user.Id = userId
	return nil
}

func (user *User) Update() (*User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	if _, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id); updateErr != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("error while trying to update user: %s", err.Error()))
	}
	
	return user, nil
}