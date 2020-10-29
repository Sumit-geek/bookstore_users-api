package app

import (
	"github.com/Sumit-geek/bookstore_users-api/controller/ping"
	"github.com/Sumit-geek/bookstore_users-api/controller/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	
	// router.GET("/users/search", controller.SearchUser)
}