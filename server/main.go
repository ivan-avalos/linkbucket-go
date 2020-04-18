/*
 *  main.go
 *  Copyright (C) 2020  Iván Ávalos <ivan.avalos.diaz@hotmail.com>
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU Affero General Public License as
 *  published by the Free Software Foundation, either version 3 of the
 *  License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU Affero General Public License for more details.
 *
 *  You should have received a copy of the GNU Affero General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"log"
	"os"

	"github.com/ivan-avalos/linkbucket/server/database"
	"github.com/ivan-avalos/linkbucket/server/setup"
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
	setup.InitValidators(e)
	setup.InitRoutes(e)

	// Start Echo HTTP server
	e.Logger.Fatal(e.Start(":" + os.Getenv("HTTP_PORT")))
}
