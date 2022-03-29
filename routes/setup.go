package routes

import (
	"github.com/chirag3003/ecommerce-golang-api/controllers"
	"github.com/gofiber/fiber/v2"
)

var conts *controllers.Controllers

func NewRoutes(cont *controllers.Controllers, app *fiber.App) {
	conts = cont
	ProductsRoutes(app.Group("/products"))
	CategoryRoutes(app.Group("/categories"))
	UserRoutes(app.Group("/user"))
	OrderRoutes(app.Group("/order"))
}
