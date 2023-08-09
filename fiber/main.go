package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, from Fiber!")
    })

  port := ":" + os.Getenv("PORT")
  log.Fatal(app.Listen(port))
}