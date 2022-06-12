package controllers

import (
	"github.com/chirag3003/ecommerce-golang-api/helpers"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"github.com/chirag3003/ecommerce-golang-api/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Me(ctx *fiber.Ctx) error
	UpdateName(ctx *fiber.Ctx) error
	AddAddress(ctx *fiber.Ctx) error
	GetAddresses(ctx *fiber.Ctx) error
	UpdateAddress(ctx *fiber.Ctx) error
	DeleteAddress(ctx *fiber.Ctx) error
}

func UserControllers() User {
	return &userRoutes{
		User: repository.NewUserRepository(conn.DB()),
	}
}

type userRoutes struct {
	User repository.UserRepository
}

func (u *userRoutes) Register(ctx *fiber.Ctx) error {
	user := &models.User{}
	err := ctx.BodyParser(user)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	if inErr := helpers.ValidateUserRegisterInput(user); inErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(inErr)
	}
	data, _ := u.User.GetUser(user.Email)
	if data != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("already exists")
	}
	if !user.CreateHash(user.Password) {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	user.SetCreateDefaults()
	res, err := u.User.Register(user)
	if err != nil {
		return err
	}
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	user.ID = id
	jwt, err := user.GetJWT()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.SendString(jwt)

}

func (u *userRoutes) Login(ctx *fiber.Ctx) error {
	user := &models.User{}
	err := ctx.BodyParser(user)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	data, _ := u.User.GetUser(user.Email)
	if data == nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	if !data.CheckPass(user.Password) {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	jwt, err := data.GetJWT()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.Status(fiber.StatusOK).SendString(jwt)
}

func (*userRoutes) Me(ctx *fiber.Ctx) error {
	data := helpers.ParseUser(ctx).Response()
	return ctx.JSON(data)
}

func (u *userRoutes) UpdateName(ctx *fiber.Ctx) error {
	user := helpers.ParseUser(ctx)
	data := &models.UpdateUserInput{}
	err := ctx.BodyParser(data)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := u.User.UpdateName(data.Name, user.ID)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.JSON(result)
}

func (u *userRoutes) AddAddress(ctx *fiber.Ctx) error {
	user := helpers.ParseUser(ctx)
	address := &models.UserAddress{}
	err := ctx.BodyParser(address)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	addAddress, err := u.User.AddAddress(user.ID, address)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(addAddress)
}

func (u *userRoutes) GetAddresses(ctx *fiber.Ctx) error {
	user := helpers.ParseUser(ctx)
	addresses, err := u.User.GetAddresses(user.ID)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.JSON(addresses)
}

func (u *userRoutes) UpdateAddress(ctx *fiber.Ctx) error {
	user := helpers.ParseUser(ctx)
	ID := ctx.Params("id")
	body := &models.UserAddressInput{}
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	response, err := u.User.UpdateAddress(user.ID, ID, body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.JSON(response)
}

func (u *userRoutes) DeleteAddress(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	result, err := u.User.DeleteAddress(ID)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.Status(fiber.StatusNoContent).JSON(result)
}
