package users

import (
	"github.com/Sumit-geek/bookstore_users-api/domain/users"
	"github.com/Sumit-geek/bookstore_users-api/service"
	"github.com/Sumit-geek/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user Id")
		c.JSON(err.Status, err)
		return
	}
	
	result, err := service.GetUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user Id")
		c.JSON(err.Status, err)
		return
	}
	
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	
	isPartial := c.Request.Method == http.MethodPatch
	user.Id = userId
	result, err := service.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	
	c.JSON(http.StatusOK, result)
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := service.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusOK, "not implemented")
}
