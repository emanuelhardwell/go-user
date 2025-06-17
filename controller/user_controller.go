package controller

import (
	"fmt"

	"github.com/emanuelhardwell/go-user/dto"
	"github.com/emanuelhardwell/go-user/service"
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	var input dto.UserCreateDTO
	if err := ctx.BodyParser(&input); err != nil {
		fmt.Printf("Error: %#v\n", err)
		return fiber.NewError(fiber.StatusBadRequest, "Verifica que tu objeto contenga: { nombre, email y password }")
	}
	out, err := service.Create(input)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.Status(fiber.StatusCreated).JSON(out)
}
