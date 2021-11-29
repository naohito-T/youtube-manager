package main

import (
	// package
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	// My subPackageを紐付ける
	"github.com/naohito-T/youtube-manager-api/src/web/routes"
)

/**
 * エラーハンドリングでログを出力する設定
 */
func init() {
	// env load
	ENV := "./env/decrypt/.env."
	ENV += os.Getenv("GO_ENV")

	err := godotenv.Load(ENV)
	if err != nil {
		logrus.Fatal("Error loading .env")
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Routes
	routes.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
