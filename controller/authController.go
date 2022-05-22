package controller

import (
	"github.com/DaisukeMatsuura/react-golang-sample/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "daisuke",
	}
	user.LastName = "matsuura"

	return c.JSON(user)
}
