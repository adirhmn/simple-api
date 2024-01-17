package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simpleapi/app"
)

func RegisterRoutes(cont *app.ServiceContainers, e *echo.Echo) {
	//init api group

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//User
	userGroup := e.Group("/v1/users")
	userHandler := UserHandler{Containers: cont}
	userGroup.Add("GET", "", userHandler.GetUsers)
	userGroup.Add("GET", "/:id", userHandler.GetUser)
	userGroup.Add("POST", "", userHandler.CreateUser)
	userGroup.Add("PUT", "/:id", userHandler.UpdateUser)
	userGroup.Add("DELETE", "/:id", userHandler.DeleteUser)
	//userGroup.Add("GET", "/:id", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Ini Detail User")
	//})
}
