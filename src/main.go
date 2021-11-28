package main

import (
	// package
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// My subPackage
	"github.com/naohito-T/youtube-manager-api/src/web/routes"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// Routes
	routes.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
