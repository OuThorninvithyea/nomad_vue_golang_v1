package handlers

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func TrustedCompany(c fiber.Ctx) error {
	return c.JSON(store.TrustedCompany)
}
