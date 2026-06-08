package user

import (
	"gotickets/internal/auth"
	"os"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	userRepository := NewRepository(db)
	jwtService := auth.NewJWTService(os.Getenv("JWT_SECRET_KEY"))
	userService := NewService(userRepository, jwtService)
	userHandler := NewHandler(userService)

	api := e.Group("api/v1/auth")
	api.POST("/register", userHandler.CreateUser)
	api.POST("/login", userHandler.LoginUser)

}

//
