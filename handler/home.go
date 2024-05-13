package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type HomeHandler struct {
	DB *sqlx.DB
}

func NewHomeHandler(db *sqlx.DB) *HomeHandler {
	return &HomeHandler{
		DB: db,
	}
}
func (h *HomeHandler) HomeHandler(c *fiber.Ctx) error {
	return c.SendString("Ledger app")
}
