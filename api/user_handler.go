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

// new ?constructor?, в общем чтобы инициализировать хендлер, видимо.
func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil { //парсит данные структуры вида params
		return err
	}
	if errors := params.Validate(); len(errors) > 0 { //???????
		return c.JSON(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	createdUser, err := h.userStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(createdUser)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var id = c.Params("id") //fetching ID from url
	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}
