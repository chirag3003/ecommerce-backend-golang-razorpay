package controllers

import (
	"github.com/chirag3003/ecommerce-golang-api/helpers"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"github.com/chirag3003/ecommerce-golang-api/repository"
	"github.com/gofiber/fiber/v2"
)

type Order interface {
	NewOrder(ctx *fiber.Ctx) error
}

func OrderControllers() Order {
	return &orderRoutes{
		Order:    repository.NewOrderRepository(conn.DB()),
		User:     repository.NewUserRepository(conn.DB()),
		Products: repository.NewProductsRepository(conn.DB()),
	}
}

type orderRoutes struct {
	Order    repository.OrderRepository
	User     repository.UserRepository
	Products repository.ProductsRepository
}

func (c *orderRoutes) NewOrder(ctx *fiber.Ctx) error {
	body := &models.NewOrderInput{}
	err := ctx.BodyParser(body)
	if err != nil || !helpers.ValidateNewOrderInput(body) {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	var products []models.OrderProduct

	for _, prod := range body.Products {
		product, err := c.Products.FindByID(prod.Product)
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		if product == nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		products = append(products, models.OrderProduct{
			Product:  *product,
			Quantity: prod.Quantity,
			Size:     prod.Size,
		})

	}

	//err = c.Order.Save(body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendStatus(fiber.StatusOK)
}
