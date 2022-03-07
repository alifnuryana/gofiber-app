package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber/database"
	"log"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Get("/", func(ctx *fiber.Ctx) error {
		err := ctx.SendString("Hello from fiber!")
		if err != nil {
			return err
		}
		return nil
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
