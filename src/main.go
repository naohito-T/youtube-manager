package main

import (
	"github.com/labstack/echo/v4"
)

// 追記します。
func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
