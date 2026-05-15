package handlers

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func Hero(c fiber.Ctx) error {
	return c.JSON(store.Hero)
}