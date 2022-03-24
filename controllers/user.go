package controllers

import (
	"github.com/chirag3003/ecommerce-golang-api/config"
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
}

func UserControllers() User {
	return &userRoutes{
		User: repository.NewUserRepository(conn.DB().Collection(config.USER_COLLECTION)),
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

	res, err := u.User.Register(user)
	if err != nil {
		return err
	}
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	user.Id = id
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
