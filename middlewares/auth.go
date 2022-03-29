package middlewares

import (
	"context"
	"github.com/chirag3003/ecommerce-golang-api/config"
	"github.com/chirag3003/ecommerce-golang-api/helpers"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsAuthenticated(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	auth, user := helpers.VerifyJWT(headers["Authorization"])
	if !auth {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	UserDB := conn.DB().Collection(config.USER_COLLECTION)
	data := &models.User{}
	ID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	err = UserDB.FindOne(context.TODO(), bson.M{"_id": ID, "email": user.Email}).Decode(data)
	if data == nil || err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	if data.UpdatedAt > user.Iat {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	ctx.Locals("user", data)
	return ctx.Next()
}

func IsAdmin(ctx *fiber.Ctx) error {
	user := helpers.ParseUser(ctx)
	for _, mail := range config.AdminUsers {
		if mail == user.Email {
			return ctx.Next()
		}
	}
	return ctx.SendStatus(fiber.StatusUnauthorized)
}
