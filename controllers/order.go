package controllers

import (
	"github.com/chirag3003/ecommerce-golang-api/models"
	"github.com/chirag3003/ecommerce-golang-api/repository"
	"github.com/gofiber/fiber/v2"
)

type Order interface {
}

func OrderControllers() Order {
	return &orderRoutes{
		Order: repository.NewOrderRepository(conn.DB()),
	}
}

type orderRoutes struct {
	Order repository.OrderRepository
}

func (c *orderRoutes) NewOrder(ctx *fiber.Ctx) error {
	body := &models.Order{}
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	body.SetCreatedAt()
	body.SetUpdatedAt()
	err = c.Order.Save(body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.JSON(body)
}
