package main

import (
	"github.com/gofiber/fiber/v2"
)

// StartAPI API for runproj
func StartAPI(){
	app := fiber.New()

	app.Get("/api/sets", func(c *fiber.Ctx) error {
		sets := GetConfigContent()
		return c.JSON(sets)
	})
	
	app.Listen(":3000")
}