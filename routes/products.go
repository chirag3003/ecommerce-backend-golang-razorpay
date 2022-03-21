package routes

import "github.com/gofiber/fiber/v2"

func ProductsRoutes(router fiber.Router) {

	router.Get("/", conts.Products.FindAll)
	router.Get("/:id", conts.Products.Find)
	router.Post("/", conts.Products.Create)
	router.Delete("/:id", conts.Products.Delete)
	router.Patch("/:id", conts.Products.Publicity)
	router.Put("/:id", conts.Products.Update)

}
