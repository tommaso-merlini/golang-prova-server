package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

   app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status": "ok",
        })
    });

    app.Static("/loaderio-04cbc2e6e8994582817d57faa8742ee5", "/loaderio-04cbc2e6e8994582817d57faa8742ee5.html")

    log.Fatal(app.Listen(":3000"))
}