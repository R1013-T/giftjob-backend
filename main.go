package main

import (
	"giftjob-backend/database"
	"giftjob-backend/routes"
)

func main() {
	database.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
