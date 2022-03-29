package controllers

import (
	"github.com/chirag3003/ecommerce-golang-api/config"
	"github.com/chirag3003/ecommerce-golang-api/helpers"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"github.com/chirag3003/ecommerce-golang-api/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Products interface {
	Create(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Publicity(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
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
	data, err := c.Products.FindAll()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func (c *productRoutes) Create(ctx *fiber.Ctx) error {
	body := &models.ProductsModel{}
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	//validating input
	inputError := helpers.ValidateProductData(body)
	if inputError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(inputError)
	}

	//setting default values for some fields
	body.SetDefaults()

	//saving the data in the database
	res, err := c.Products.Save(body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	//fetching the object id of the created product
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	body.ID = id

	return ctx.Status(fiber.StatusOK).JSON(body)
}

func (c *productRoutes) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := c.Products.Find(id)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func (c *productRoutes) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	result, err := c.Products.Delete(id)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	return ctx.Status(fiber.StatusNoContent).JSON(result)
}

func (c *productRoutes) Publicity(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	type inputStruct struct {
		Public bool `json:"public"`
	}
	input := &inputStruct{}
	err := ctx.BodyParser(input)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	find, err := c.Products.Find(id)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	if find.Stock == 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("Stock Too Low")
	}

	visibility, err := c.Products.ChangeVisibility(id, input.Public)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(visibility)
}

func (c *productRoutes) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	body := &models.ProductsModel{}
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	//validating input
	inputError := helpers.ValidateProductData(body)
	if inputError != nil {

		log.Println(inputError)
		return ctx.Status(fiber.StatusBadRequest).JSON(inputError)
	}

	//saving the data in the database
	res, err := c.Products.Update(id, body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	//fetching the object id of the created product

	return ctx.Status(fiber.StatusOK).JSON(res)
}
