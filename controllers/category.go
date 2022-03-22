package controllers

import (
	"github.com/chirag3003/ecommerce-golang-api/config"
	"github.com/chirag3003/ecommerce-golang-api/helpers"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"github.com/chirag3003/ecommerce-golang-api/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Categories interface {
	FindAll(ctx *fiber.Ctx) error
	FindCategory(ctx *fiber.Ctx) error
	FindSubcategory(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	CreateSubcategory(ctx *fiber.Ctx) error
	UpdateCategory(ctx *fiber.Ctx) error
	UpdateSubcategory(ctx *fiber.Ctx) error
	DeleteCategory(ctx *fiber.Ctx) error
	DeleteSubcategory(ctx *fiber.Ctx) error
	Visibility(ctx *fiber.Ctx) error
}

type categoryRoutes struct {
	Category repository.CategoryRepository
}

func CategoryControllers() Categories {
	return &categoryRoutes{
		Category: repository.NewCategoryRepo(conn.DB().Collection(config.CATEGORIES_COLLECTION)),
	}
}

func (c *categoryRoutes) FindAll(ctx *fiber.Ctx) error {
	data, err := c.Category.FindAll()
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.JSON(data)
}
func (c *categoryRoutes) Create(ctx *fiber.Ctx) error {
	body := &models.Category{}
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	//validating input
	inputError := helpers.ValidateCategory(body)
	if inputError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(inputError)
	}

	//setting default values for some fields
	body.SetDefaults()
	body.Subcategories = []models.Subcategory{}

	//saving the data in the database
	res, err := c.Category.SaveCategory(body)
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
func (c *categoryRoutes) CreateSubcategory(ctx *fiber.Ctx) error {
	body := &models.Subcategory{}
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	categoryID := ctx.Params("id")
	body.ID = primitive.NewObjectID()
	body.SetDefaults()
	data, err := c.Category.SaveSubcategory(body, categoryID)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.Status(200).JSON(data)
}
func (c *categoryRoutes) FindCategory(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	data, err := c.Category.FindCategory(ID)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	return ctx.JSON(data)
}
func (c *categoryRoutes) FindSubcategory(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	data, err := c.Category.FindSubcategory(ID)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	return ctx.JSON(data)
}
func (c *categoryRoutes) UpdateCategory(ctx *fiber.Ctx) error {
	body := &models.CategoryUpdateInput{}
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	ID := ctx.Params("id")
	data, err := c.Category.UpdateCategory(ID, body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	return ctx.JSON(data)
}
func (c *categoryRoutes) UpdateSubcategory(ctx *fiber.Ctx) error {
	body := &models.Subcategory{}
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	ID := ctx.Params("id")
	data, err := c.Category.UpdateSubcategory(ID, body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	return ctx.JSON(data)
}
func (c *categoryRoutes) DeleteCategory(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	data, err := c.Category.DeleteCategory(ID)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	return ctx.Status(fiber.StatusNoContent).JSON(data)
}
func (c *categoryRoutes) DeleteSubcategory(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	data, err := c.Category.DeleteSubcategory(ID)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	return ctx.Status(fiber.StatusNoContent).JSON(data)
}
func (c *categoryRoutes) Visibility(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	body := &struct {
		Public bool `json:"public"`
	}{}
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.SendStatus(401)
	}
	visibility, err := c.Category.ChangeVisibility(ID, body.Public)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	return ctx.Status(fiber.StatusOK).JSON(visibility)
}
