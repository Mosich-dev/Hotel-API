package API

import (
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

func (h *UserHandler) HandleInsertUser(ctx *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) != 0 {
		return ctx.JSON(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	createdUser, err := h.userStore.InsertUser(ctx.Context(), user)
	if err != nil {
		return err
	}

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

func (h *UserHandler) HandlePutUsers(ctx *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) HandleDeleteUser(ctx *fiber.Ctx) error {
	type unMarshalUserID struct {
		Id string
	}
	var userID unMarshalUserID

	if err := ctx.BodyParser(&userID); err != nil {
		return err
	}
	if err := h.userStore.DeleteUser(ctx.Context(), userID.Id); err != nil {
		return err
	}
	return ctx.JSON(userID)
}
