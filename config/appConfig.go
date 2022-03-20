package config

import (
	"github.com/chirag3003/ecommerce-golang-api/db"
	"github.com/gofiber/fiber/v2"
)

type AppConfig struct {
	Conn db.Connection
	App  *fiber.App
}
