package main

import (
	"github.com/chirag3003/ecommerce-golang-api/config"
	"github.com/chirag3003/ecommerce-golang-api/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

var appConfig *config.AppConfig

func main() {
	//Loading Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//connecting mongo
	client := db.ConnectMongo()
	defer client.Close()

	//creating fiber app
	app := fiber.New()

	//setting up configs
	appConfig.Conn = client
	appConfig.App = app

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
