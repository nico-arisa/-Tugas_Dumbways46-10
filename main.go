package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	//routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	fmt.Println("Server Started on Port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))

}
