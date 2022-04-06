package controllers

import "github.com/gofiber/fiber/v2"

type Images interface {
	Upload(ctx *fiber.Ctx) error
}

func ImagesControllers() Images {
	return &imagesRoutes{}
}

type imagesRoutes struct {
}

func (i imagesRoutes) Upload(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}
