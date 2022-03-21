package controllers

import (
	"github.com/chirag3003/ecommerce-golang-api/config"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"github.com/chirag3003/ecommerce-golang-api/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Products interface {
	Create(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type productRoutes struct {
	Products repository.ProductsRepository
}

func ProductsControllers() Products {
	return &productRoutes{
		Products: repository.NewProductsRepository(conn.DB().Collection(config.PRODUCTS_COLLECTION)),
	}
}

func (c *productRoutes) FindAll(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}

func (c *productRoutes) Create(ctx *fiber.Ctx) error {
	body := &models.ProductsModel{
		Title:       "Product",
		Description: "Description",
	}
	body.SetDefaults()
	res, err := c.Products.Save(body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	body.ID = id

	return ctx.Status(fiber.StatusOK).JSON(body)
}

func (c *productRoutes) Find(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}

func (c *productRoutes) Delete(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}
