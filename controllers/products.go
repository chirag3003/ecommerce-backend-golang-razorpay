package controllers

import (
	"fmt"
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
	GetCartData(ctx *fiber.Ctx) error
	GetStockExcel(ctx *fiber.Ctx) error
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
	if data == nil {
		data = []models.ProductsModel{}
	}
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func (c *productRoutes) Create(ctx *fiber.Ctx) error {
	body := &models.ProductsModel{}
	err := ctx.BodyParser(body)
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	//validating input
	inputError := helpers.ValidateProductData(body)
	if inputError != nil {
		fmt.Println(inputError)
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
	slug := ctx.Params("slug")
	data, err := c.Products.Find(slug)
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
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	_, err = c.Products.FindByID(ID)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
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

	//validating slug
	find, _ := c.Products.Find(body.Slug)
	if find != nil && find.ID.String() != body.ID.String() {
		return ctx.Status(fiber.StatusBadRequest).SendString("slug already in use")
	}

	//saving the data in the database
	res, err := c.Products.Update(id, body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *productRoutes) GetCartData(ctx *fiber.Ctx) error {
	var body []struct {
		ID       primitive.ObjectID `json:"id"`
		Size     string             `json:"size"`
		Quantity int                `json:"quantity"`
	}
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}
	var resp []models.CartSearchResult
	for _, product := range body {
		data, err := c.Products.FindByID(product.ID)
		if err != nil {
			continue
		}
		var stock int
		for _, size := range data.Sizes {
			if size.Name == product.Size {
				stock = size.Stock
			}
		}
		resp = append(resp, models.CartSearchResult{
			Product:  data,
			Size:     product.Size,
			Stock:    stock,
			Quantity: product.Quantity,
		})

	}
	fmt.Println(resp)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *productRoutes) GetStockExcel(ctx *fiber.Ctx) error {
	// Getting Product Data
	data, err := c.Products.FindAll()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	//Generating Excel File
	file, err := helpers.GenerateStockExcel(data, ctx)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	return ctx.SendStream(file)
}
