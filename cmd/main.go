package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

type User struct {
	gorm.Model
	Name     string `json:"name"     validate:"required"       gorm:"type:varchar(100);not null"`
	Email    string `json:"email"    validate:"required,email"  gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `json:"password"  validate:"required,min=6"  gorm:"type:varchar(255);not null"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{})

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	e := echo.New()
	// e.Use(middleware.RequestLogger())

	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", func(c *echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return err
		}
		if err = c.Validate(u); err != nil {
			return err
		}
		result := db.Create(u)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": result.Error.Error(),
			})
		}
		return c.JSON(http.StatusOK, u)
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
