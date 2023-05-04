package main

import (
	"fmt"
	"log"
	"server/db"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
)

type Product struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
}

func main() {
	app := fiber.New()
	db.Init()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Get("/product/:id", func(c *fiber.Ctx) error {
		product := &Product{}
		coll := mgm.Coll(product)

		// Find and decode the doc to a book model.
		err := coll.FindByID(c.Params("id"), product)

		if err != nil {
			fmt.Println(err)
		}

		return c.JSON(product)
	})

	app.Static("/loaderio-04cbc2e6e8994582817d57faa8742ee5", "/loaderio-04cbc2e6e8994582817d57faa8742ee5.html")

	log.Fatal(app.Listen(":3000"))
}