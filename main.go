package main

import (
	"giftjob-backend/database"
	"giftjob-backend/routes"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.Init()
	e := routes.Init()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":8080"))
}
