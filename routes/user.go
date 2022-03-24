package routes

import "github.com/gofiber/fiber/v2"

func UserRoutes(router fiber.Router) {
	router.Post("/login", conts.User.Login)
	router.Post("/register", conts.User.Register)
	router.Get("/me", conts.User.Me)
}
