package adapters

import (
	"fmt"
	"kington/core"

	"github.com/gofiber/fiber/v2"
)

type HttpOrderHandler struct {
	service core.OrderService
}

func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order core.Order
	if err := c.BodyParser(&order); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.service.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(order)
}
