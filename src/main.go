package main

import (
	// package
	"github.com/labstack/echo/v4"
	// my subPackage
	"github.com/naohito-T/youtube-manager-api/src/web/routes"
)

// 追記します。
func main() {
	e := echo.New()

	// Routes
	routes.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
