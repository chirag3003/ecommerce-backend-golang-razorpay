package routes

import "github.com/gofiber/fiber/v2"

func OrderRoutes(router fiber.Router) {
	router.Post("/", conts.Order.NewOrder)
}
