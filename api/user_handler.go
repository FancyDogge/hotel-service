package api

import (
	"github.com/FancyDogge/hotel-service/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Kekw",
		LastName:  ":O",
	}
	return c.JSON(u)
}

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("James")
}
