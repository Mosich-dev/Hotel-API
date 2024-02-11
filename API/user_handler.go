package API

import (
	"errors"
	"github.com/Mosich-dev/Hotel-API/db"
	"github.com/Mosich-dev/Hotel-API/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
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
	var params types.InsertUserParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	if valErrs := params.ValidateAll(); len(valErrs) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(valErrs)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	createdUser, err := h.userStore.InsertUser(ctx.Context(), user)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(createdUser)
}

func (h *UserHandler) HandleGetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := h.userStore.GetUserByID(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.JSON(map[string]string{"error": "not found"})
		}
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

func (h *UserHandler) HandlePutUser(ctx *fiber.Ctx) error {
	var (
		params types.UpdateUserParams
		userID = ctx.Params("id")
	)

	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	if err := h.userStore.UpdateUser(ctx.Context(), userID, params); err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(map[string]string{"updated": userID})
}
