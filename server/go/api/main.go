package main

import (
	"Slack/apifunc"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	log.Println(i)
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewServer() *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// setting static files
	e.Static("/static/img", "./static/img")
	e.Static("/static/css", "./static/css")
	e.Static("/static/js", "./static/js")
	e.Static("/static/font", "./static/font")

	// setting middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Fprintf(os.Stderr, "Request: %v\n", string(reqBody))
		// fmt.Fprintf(os.Stderr, "Header: %v\n", c.Request().Header)
	}))

	// setting routes
	requiredAuth := e.Group("")

	requiredAuth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		// TokenLookup: "header: Authorization,cookie:token",
	}))

	// api routes
	// http://localhost:8080/user : GET apifunc->user.go->UserGet()
	requiredAuth.GET("/user", apifunc.UserGet)
	// http://localhost:8080/user : POST apifunc->user.go->UserPost()
	e.POST("/user", apifunc.UserPost)
	// http://localhost:8080/user : PUT apifunc->user.go->UserPut()
	requiredAuth.PUT("/user", apifunc.UserPut)
	requiredAuth.DELETE("/user", apifunc.UserDelete)
	// http://localhost:8080/login : POST apifunc->login.go->LoginPost()
	e.POST("/login", apifunc.LoginPost)
	return e
}

func main() {
	e := NewServer()
	e.Logger.Fatal(e.Start(":8080"))
}
