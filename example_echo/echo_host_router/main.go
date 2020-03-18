package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	lg := e.Host("localhost:1234")
	lg.GET("/test", test1)

	ig := e.Host("127.0.0.1:1234")
	ig.GET("/test", test2)

	e.GET("/test", test3)

	e.HideBanner = true

	for _, r := range e.Routes() {
		fmt.Printf("%#v", r)
	}

	if err := e.Start(":1234"); err != nil {
		log.Fatal(err)
	}
}

func test1(c echo.Context) error {
	return c.String(200, "test router for host \"localhost\"")
}

func test2(c echo.Context) error {
	return c.String(200, "test router for host \"127.0.0.1\"")
}

func test3(c echo.Context) error {
	return c.String(200, "test router for default")
}
