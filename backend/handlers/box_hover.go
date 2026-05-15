package handlers

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func BoxHover(c fiber.Ctx) error {
	return c.JSON(store.BoxHover)
}