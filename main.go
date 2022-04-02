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
	db 				*gorm.DB						= config.SetUpDatabaseConnection()
	userRepo		repository.UserRepo				= repository.NewUserRepo(db)
	spRepo			repository.SmartPhoneRepo		= repository.NewSmartPhoneRepo(db)

	jwtService		service.JWTService				= service.NewJWTService()
	authService		service.AuthService				= service.NewAuthService(userRepo)
	spService		service.SmartPhoneService		= service.NewSmartPhoneService(spRepo)

	authController 	controller.AuthController		= controller.NewAuthController(authService, jwtService)
	spController	controller.SmartPhoneController	= controller.NewSmartPhoneController(spService) 
)
func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	spRoutes := r.Group("api/smartPhone")
	{
		spRoutes.GET("", spController.SpGetAll)
		spRoutes.GET("/:condition", spController.SpGetByCond)
		spRoutes.POST("", spController.SpPost)
		spRoutes.PUT("/:id", spController.SpPut)
		spRoutes.DELETE("/:id", spController.SpDelete)
	}
	r.Run(":8888")
}