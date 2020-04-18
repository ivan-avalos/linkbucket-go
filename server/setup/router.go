/*
 *  router.go
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

package setup

import (
	"os"

	"github.com/ivan-avalos/linkbucket-go/server/controllers"
	"github.com/ivan-avalos/linkbucket-go/server/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitRoutes initializes routes
func InitRoutes(e *echo.Echo) {
	e.POST("/api/register", controllers.CreateUser)
	e.POST("/api/token", controllers.Authenticate)
	auth := e.Group("/api")
	auth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &utils.Token{},
		SigningKey: []byte(os.Getenv("TOKEN_PASSWORD")),
	}))
	{
		auth.GET("/user", controllers.GetUser)
		auth.PUT("/user", controllers.UpdateUser)
		auth.DELETE("/user", controllers.DeleteUser)

		auth.GET("/link", controllers.GetLinks)
		auth.GET("/link/:id", controllers.GetLink)
		auth.POST("/link", controllers.CreateLink)
		auth.PUT("/link/:id", controllers.UpdateLink)
		auth.DELETE("/link/:id", controllers.DeleteLink)

		auth.GET("/tag", controllers.GetTags)
		auth.GET("/tag/:slug", controllers.GetLinksForTag)
		auth.GET("/search", controllers.GetLinksForSearch)
	}
}
