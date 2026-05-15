package handlers

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func AdTwo(c fiber.Ctx) error {
	return c.JSON(store.AdTwo)
}