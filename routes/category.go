package routes

import (
	"github.com/chirag3003/ecommerce-golang-api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(router fiber.Router) {
	router.Get("/", conts.Categories.FindAll)
	router.Get("/subcategory/:id", conts.Categories.FindSubcategory)
	router.Get("/category/:id", conts.Categories.FindCategory)
	router.Post("/", middlewares.IsAuthenticated, middlewares.IsAdmin, conts.Categories.Create)
	router.Post("/subcategory/:id", middlewares.IsAuthenticated, middlewares.IsAdmin, conts.Categories.CreateSubcategory)
	router.Put("/category/:id", middlewares.IsAuthenticated, middlewares.IsAdmin, conts.Categories.UpdateCategory)
	router.Put("/subcategory/:id", middlewares.IsAuthenticated, middlewares.IsAdmin, conts.Categories.UpdateSubcategory)
	router.Delete("/category/:id", middlewares.IsAuthenticated, middlewares.IsAdmin, conts.Categories.DeleteCategory)
	router.Delete("/subcategory/:id", middlewares.IsAuthenticated, middlewares.IsAdmin, conts.Categories.DeleteSubcategory)
	router.Patch("/:id", middlewares.IsAuthenticated, middlewares.IsAdmin, conts.Categories.Visibility)
}
