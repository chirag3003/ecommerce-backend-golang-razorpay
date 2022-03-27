package routes

import (
	"github.com/chirag3003/ecommerce-golang-api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	router.Post("/login", conts.User.Login)
	router.Post("/register", conts.User.Register)
	router.Post("/address", middlewares.IsAuthenticated, conts.User.AddAddress)
	router.Get("/me", middlewares.IsAuthenticated, conts.User.Me)
	router.Get("/address", middlewares.IsAuthenticated, conts.User.GetAddresses)
	router.Patch("/name", middlewares.IsAuthenticated, conts.User.UpdateName)
}
