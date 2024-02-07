package api

import (
	"github.com/FancyDogge/hotel-service/db"
	"github.com/FancyDogge/hotel-service/types"
	"github.com/gofiber/fiber/v2"
)

// handler, чтобы работать с юзером и db
type UserHandler struct {
	userStore db.UserStore //для возможности юзать разные интерфейсы
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id") //fetching ID from url
	user, err := h.userStore.GetUserByID(id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Kekw",
		LastName:  ":O",
	}
	return c.JSON(u)
}
