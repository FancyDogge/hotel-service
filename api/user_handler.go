package api

import (
	"errors"

	"github.com/FancyDogge/hotel-service/db"
	"github.com/FancyDogge/hotel-service/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	if errors := params.ValidateCreateUser(); len(errors) > 0 { //???????
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

// c.ParamsParser() nujno razuznat'
func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var (
		id     = c.Params("id")
		params types.UpdateUserParams
	)
	oid, err := primitive.ObjectIDFromHex(id) //разобраться почему если конвертинг типов ниже BodyParser, то ничего не работает
	if err != nil {
		return err
	}
	if err := c.BodyParser(&params); err != nil { //parse everything into these &params / all params are passed through fiber.Ctx vrode??
		return err
	}
	// if errors := params.ValidateCreateUser(); len(errors) > 0 { //???????
	// 	return c.JSON(errors)
	// }
	filter := bson.M{"_id": oid}
	if err := h.userStore.UpdateUser(c.Context(), filter, params); err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": id})
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var id = c.Params("id") //fetching ID from url
	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"message": "not found"})
		}
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

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	var id = c.Params("id")
	res, err := h.userStore.DeleteUser(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"message": "not found"})
		}
		return err
	}
	return c.JSON(res)
}
