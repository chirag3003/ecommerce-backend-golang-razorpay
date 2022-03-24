package main

import (
	"github.com/chirag3003/ecommerce-golang-api/controllers"
	"github.com/chirag3003/ecommerce-golang-api/db"
	"github.com/chirag3003/ecommerce-golang-api/helpers"
	"github.com/chirag3003/ecommerce-golang-api/middlewares"
	"github.com/chirag3003/ecommerce-golang-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	//Loading Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	helpers.VerifyENV()

	//connecting mongo
	client := db.ConnectMongo()
	defer client.Close()

	//creating fiber app
	app := fiber.New()
	app.Use(logger.New())

	//Setting Up Controllers
	middlewares.SetupMiddlewares(client)
	routes.NewRoutes(controllers.NewControllers(client), app)
	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
