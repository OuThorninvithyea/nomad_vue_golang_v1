package handlers

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func HoverCard(c fiber.Ctx) error {
	return c.JSON(store.HoverCards)
}