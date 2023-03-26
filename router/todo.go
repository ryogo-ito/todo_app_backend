package router

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"todoappbackend/model"
)

func GetTodosHandler(c echo.Context) error {
	todos, err := model.GetTodos()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, todos)
}

type ReqTodo struct {
	Name string `json:"name"`
}

func AddTodoHandler(c echo.Context) error {
	var req ReqTodo

	err := c.Bind(&req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var todo *model.Todo

	todo, err = model.AddTodo(req.Name)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, todo)
}

func ChangeCompletedTodoHandler(c echo.Context) error {
	todoID, err := uuid.Parse(c.Param("todoID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	err = model.ChangeCompletedTodo(todoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent(http.StatusNoContent)
}

func DeleteTodoHandler(c echo.Context) error {
	todoID, err := uuid.Parse(c.Param("todoID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	err = model.DeleteTodo(todoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.NoContent(http.StatusOK)
}
