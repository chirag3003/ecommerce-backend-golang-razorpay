package routes

import (
	"github.com/chirag3003/ecommerce-golang-api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(router fiber.Router) {
	router.Post("/", middlewares.IsAuthenticated, conts.Order.NewOrder)
	router.Post("/razorpay/paid", conts.Order.OrderPaid)
	router.Post("/razorpay/event", conts.Order.OrderEvents)
}
