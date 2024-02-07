package api

import (
	"context"

	"github.com/FancyDogge/hotel-service/db"
	"github.com/FancyDogge/hotel-service/types"
	"github.com/gofiber/fiber/v2"
)

// handler, чтобы работать с юзером и db
type UserHandler struct {
	userStore db.UserStore //для возможности юзать разные интерфейсы
}

// new ?constructor?, в общем чтобы инициализировать хендлер, видимо.
func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id") //fetching ID from url
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)
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
