package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nazartymiv/jwt-auth/controllers"
	"github.com/nazartymiv/jwt-auth/initializers"
	"github.com/nazartymiv/jwt-auth/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DbConnection()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}