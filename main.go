package main

import (
	"github.com/labstack/echo/v4"
	"todoappbackend/model"
	"todoappbackend/router"
)

func main() {
	sqlDB := model.DBConnection()
	defer sqlDB.Close()
	e := echo.New()
	router.SetRouter(e)
}
