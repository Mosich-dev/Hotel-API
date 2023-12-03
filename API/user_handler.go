package API

import (
	"github.com/Mosich-dev/Hotel-API/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(ctx *fiber.Ctx) error {
	u := types.User{
		ID:        "7",
		FirstName: "mostafa",
		LastName:  "chegeni",
	}
	return ctx.JSON(u)
}

func HandleGetUser(ctx *fiber.Ctx) error {
	return ctx.JSON("Hi! id")
}
