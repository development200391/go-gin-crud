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
	config.DB.AutoMigrate(&model.User{}, &model.Role{}, &model.Permission{}, &model.RolePermission{})

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	auth := r.Group("/", middleware.AuthMiddleware())
	{
		auth.POST("/users", handler.CreateUser)
		auth.GET("/users", handler.GetUsers)
		auth.GET("/users/:id", handler.GetUserByID)
		auth.PUT("/users/:id", handler.UpdateUser)
		auth.DELETE("/users/:id", handler.DeleteUser)

		auth.POST("/roles", handler.CreateRole)
		auth.GET("/roles", handler.GetRoles)
		auth.GET("/roles/:id", handler.GetRoleByID)
		auth.PUT("/roles/:id", handler.UpdateRole)
		auth.DELETE("/roles/:id", handler.DeleteRole)

		auth.POST("/permissions", handler.CreatePermission)
		auth.GET("/permissions", handler.GetPermissions)
		auth.GET("/permissions/:id", handler.GetPermissionByID)
		auth.PUT("/permissions/:id", handler.UpdatePermission)
		auth.DELETE("/permissions/:id", handler.DeletePermission)

		auth.POST("/role-permissions", handler.CreateRolePermission)
		auth.GET("/role-permissions", handler.GetRolePermissions)
		auth.GET("/role-permissions/:role_id/:perm_id", handler.GetRolePermissionByID)
		auth.PUT("/role-permissions/:role_id/:perm_id", handler.UpdateRolePermission)
		auth.DELETE("/role-permissions/:role_id/:perm_id", handler.DeleteRolePermission)
	}

	r.Run(":8080")
}
