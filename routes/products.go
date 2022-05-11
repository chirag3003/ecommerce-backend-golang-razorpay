package routes

import (
	"github.com/chirag3003/ecommerce-golang-api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ProductsRoutes(router fiber.Router) {

	router.Get("/", conts.Products.FindAll)
	router.Get("/:slug", conts.Products.Find)
	router.Get("/stocks/excel", conts.Products.GetStockExcel)
	router.Post("/", middlewares.IsAuthenticated, middlewares.IsAdmin, conts.Products.Create)
	router.Post("/cartData", conts.Products.GetCartData)
	router.Delete("/:id", middlewares.IsAuthenticated, middlewares.IsAdmin, conts.Products.Delete)
	router.Patch("/:id", middlewares.IsAuthenticated, middlewares.IsAdmin, conts.Products.Publicity)
	router.Put("/:id", middlewares.IsAuthenticated, middlewares.IsAdmin, conts.Products.Update)

}
