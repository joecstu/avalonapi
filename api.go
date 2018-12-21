package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type ID struct {
	Username string
	Password string
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("login",Login)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {

	return c.JSON(http.StatusOK, "Hello, World!")
}
func Login(c echo.Context) error {

	var request struct {
		Username string
		Password string
	}
	var response struct {
		Status        string `json:",omitempty"` //"success | error | inactive"
		StatusMessage string `json:",omitempty"`
	}

	c.Bind(&request)

	//fmt.Println(request.Username)
	//fmt.Println(request.Password)

	fmt.Println(request)


	if strings.Compare(request.Username, "test") == 0 && strings.Compare(request.Password, "1234") == 0 {
		response.Status="เข้าสู่ระบบได้จ้า"
		response.StatusMessage="เข้าเกมได้เลย"
	}else{
		response.Status="เข้าสู่ระบบไม่ได้จ้า"
		response.StatusMessage="ไปส่องกะหรี่"
	}

	return c.JSON(http.StatusOK,response)
}