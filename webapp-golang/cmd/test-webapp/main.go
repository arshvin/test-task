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
	serverAddr = kingpin.Flag("address", "IP Address to bind to").Default("0.0.0.0").String()
	serverPort = kingpin.Flag("port", "Port for listening").Default("8080").String()
)

func main() {
	kingpin.Parse()

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, e.Routes(),"  ")
	}).Name = "available-routes"

	e.GET("/cpu", func(c echo.Context) error {
		return handlers.CpuAmount(c)
	}).Name = "cpu-amount"

	e.GET("/ip", func(c echo.Context) error {
		return handlers.AddrIpV4Amount(c)
	}).Name = "ip-addr-amount"

	e.GET("/date", func(c echo.Context) error {
		return handlers.CurrentDateTime(c, true)
	}).Name = "date-now"

	e.GET("/time", func(c echo.Context) error {
		return handlers.CurrentDateTime(c, false)
	}).Name = "time-amount"

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", *serverAddr, *serverPort)))

}
