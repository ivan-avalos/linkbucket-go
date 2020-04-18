package main

import (
	"os"

	"github.com/ivan-avalos/linkbucket/controllers"
	"github.com/ivan-avalos/linkbucket/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func initRoutes(e *echo.Echo) {
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