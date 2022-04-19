package routes

import "github.com/gofiber/fiber/v2"

func ImagesRoutes(router fiber.Router) {
	router.Post("/", conts.Images.Upload)
}
