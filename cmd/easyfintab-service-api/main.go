package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "web/static")

	e.HideBanner = true
	e.HidePort = true

	e.Logger.Fatal(e.Start(":8080"))
}
