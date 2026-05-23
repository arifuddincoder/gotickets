package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
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
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func main() {

	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
		return c.JSON(http.StatusOK, u)
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
