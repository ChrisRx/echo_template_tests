package main

//go:generate qtc -dir ./templates -ext html

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	"./templates"
)

func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, templates.Hello("World"))
	}
}

func main() {
	e := echo.New()
	e.Debug()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Get("/", hello())
	e.Run(standard.New(":3000"))
}
