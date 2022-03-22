package routes

import "github.com/gofiber/fiber/v2"

func CategoryRoutes(router fiber.Router) {
	router.Get("/", conts.Categories.FindAll)
	router.Get("/subcategory/:id", conts.Categories.FindSubcategory)
	router.Get("/category/:id", conts.Categories.FindCategory)
	router.Post("/", conts.Categories.Create)
	router.Post("/subcategory/:id", conts.Categories.CreateSubcategory)
	router.Put("/category/:id", conts.Categories.UpdateCategory)
	router.Put("/subcategory/:id", conts.Categories.UpdateSubcategory)
	router.Delete("/category/:id", conts.Categories.DeleteCategory)
	router.Delete("/subcategory/:id", conts.Categories.DeleteSubcategory)
	router.Patch("/:id", conts.Categories.Visibility)
}
