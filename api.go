package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// StartAPI API for runproj
func StartAPI(){
	app := fiber.New()

	app.Get("/api/sets", func(c *fiber.Ctx) error {
		sets := GetConfigContent()
		jsonSets, err := json.Marshal(sets)

		if err != nil {
			return err
		}

		c.Response().Header.Add("Content-Type", "application/json")

		return c.SendString(string(jsonSets))
	})
	
	app.Listen(":3000")
}