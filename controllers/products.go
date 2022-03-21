package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Products interface {
	Create(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type productRoutes struct {
	Products *mongo.Collection
}

func ProductRoutes() Products {
	return &productRoutes{
		Products: conn.DB().Collection("products"),
	}
}

func (c *productRoutes) FindAll(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}

func (c *productRoutes) Create(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}

func (c *productRoutes) Find(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}

func (c *productRoutes) Delete(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}
