package handlers

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func SearchDropdown(c fiber.Ctx) error {
	return c.JSON(store.SearchDropdown)
}
