package main

import (
	"github.com/DaisukeMatsuura/react-golang-sample/database"
	"github.com/DaisukeMatsuura/react-golang-sample/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
