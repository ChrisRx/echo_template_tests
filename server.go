package main

//go:generate qtc -dir templates -ext html

import (
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	"./templates"
)

var hashKey = []byte("secretkeygoeshere")
var s = securecookie.New(hashKey, nil)

func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, templates.Hello("World"))
	}
}

func Auth() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) error {
			cookies := c.Request().Header().Get("Cookie")
			if cookies == "" {
				encoded, err := s.Encode("testcookie", "testvalue")
				if err != nil {
					log.Println(err)
				}
				n := &http.Cookie{
					Name:  "testcookie",
					Value: encoded,
					Path:  "/",
				}
				cookies = n.String()
				c.Response().Header().Add("Set-Cookie", cookies)
			}
			log.Printf("Cookies: %v\n", cookies)
			return next.Handle(c)
		})
	}
}

func main() {
	e := echo.New()
	e.Debug()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(Auth())
	e.Get("/", hello())
	e.Run(standard.New(":3000"))
}
