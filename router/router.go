package router

import (
	_ "net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouter(e *echo.Echo) error {

	// 諸々の設定(*1)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/todos", GetTodosHandler)
	e.POST("api/todos", AddTodoHandler)
	e.PUT("api/todos/:todoID", ChangeCompletedTodoHandler)
	e.DELETE("/api/todos/:todoID", DeleteTodoHandler)

	err := e.Start(":8000")
	return err
}
