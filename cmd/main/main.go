package main

import (
	"net/http"

	"github.com/benjaminbrassart/test-go-frontend/frontend"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(frontend.GetStaticFilesRoot()),
	}))

	if err := app.Listen("127.0.0.1:3000"); err != nil {
		logrus.WithError(err).Fatal("failed to start http app")
	}
}
