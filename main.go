package main

import (
	"go-api/config"
	"go-api/controller"
	"go-api/repository"
	"go-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
var(
	db 				*gorm.DB					= config.SetUpDatabaseConnection()
	userRepo		repository.UserRepo			= repository.NewUserRepo(db)
	jwtService		service.JWTService			= service.NewJWTService()
	authService		service.AuthService			= service.NewAuthService(userRepo)
	authController 	controller.AuthController	= controller.NewAuthController(authService, jwtService)
)
func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run(":8888")
}