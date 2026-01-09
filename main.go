package main

import (
	"go-gin-crud/config"
	"go-gin-crud/handler"
	"go-gin-crud/middleware"
	"go-gin-crud/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()
	config.DB.AutoMigrate(&model.User{})

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	auth := r.Group("/", middleware.AuthMiddleware())
	{
		auth.GET("/users", handler.GetUsers)
		auth.GET("/users/:id", handler.GetUserByID)
		auth.PUT("/users/:id", handler.UpdateUser)
		auth.DELETE("/users/:id", handler.DeleteUser)
	}

	r.Run(":8080")
}
