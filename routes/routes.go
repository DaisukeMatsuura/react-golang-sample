package routes

import (
	"github.com/DaisukeMatsuura/react-golang-sample/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}
