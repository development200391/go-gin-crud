package main

import (
	"go-gin-crud/config"
	"go-gin-crud/handler"
	"go-gin-crud/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// connect DB
	config.ConnectDB()

	// auto migrate
	config.DB.AutoMigrate(&model.User{})

	r.GET("/users", handler.GetUsers)
	r.GET("/users/:id", handler.GetUserByID)
	r.POST("/users", handler.CreateUser)
	r.PUT("/users/:id", handler.UpdateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	r.Run(":8080")
}
