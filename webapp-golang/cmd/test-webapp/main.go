package main

import (
	"fmt"
	"net/http"

	handlers "test-webapp/internal/handlers"

	"github.com/alecthomas/kingpin/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	serverAddr = kingpin.Flag("a", "IP Address to bind to").Default("0.0.0.0").String()
	serverPort = kingpin.Arg("p", "Port for listening").Default("8080").String()
)

func main() {
	kingpin.Parse()

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/cpu", func(c echo.Context) error {
		return handlers.CpuAmount(c)
	})

	e.GET("/ip", func(c echo.Context) error {
		return handlers.AddrIpV4Amount(c)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", *serverAddr, *serverPort)))

}
