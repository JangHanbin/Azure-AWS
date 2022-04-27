package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	controller "mini-contents-hub/internal/controller"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	controller.StorageController(app)

	log.Fatal(app.Listen(":3000"))
}
