package main

import (
	"net/http"

	"jwt-go/controllers"
	"jwt-go/initializers"
	"jwt-go/middleware"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	e := echo.New()

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/signup", controllers.Signup)
	e.POST("/login", controllers.Login)
	e.POST("/logout", controllers.Logout)
	e.GET("/validate", controllers.Validate, middleware.RequireAuth)

	e.Logger.Fatal(e.Start(":3000"))
}
