package main

import (
	"log"
	"os"

	"github.com/ivan-avalos/linkbucket/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Load config from .env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Initialise DB
	database.Init()

	// Initialise Echo HTTP server
	e := echo.New()
	e.Use(middleware.Logger())
	initValidators(e)
	initRoutes(e)

	// Start Echo HTTP server
	e.Logger.Fatal(e.Start(":" + os.Getenv("HTTP_PORT")))
}
