package API

import (
	"fmt"
	"github.com/Mosich-dev/Hotel-API/db"
	"github.com/Mosich-dev/Hotel-API/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandlePostUser(ctx *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	if err := params.Validate(); err != nil {
		return err
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	createdUser, err := h.userStore.InsertUser(ctx.Context(), user)
	if err != nil {
		return err
	}
	fmt.Println(createdUser)

	return ctx.JSON(createdUser)
}

func (h *UserHandler) HandleGetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := h.userStore.GetUserByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.JSON(user)
}

func (h *UserHandler) HandleGetUsers(ctx *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(users)
}
