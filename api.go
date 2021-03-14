package main

import (
	"github.com/gofiber/fiber/v2"
)

func main(){
	StartAPI()
	// args := os.Args[1:]

	// sets := GetConfigContent()

	// ExecuteSelectedSets(sets, args
}

// StartAPI API for runproj
func StartAPI(){
	app := fiber.New()

	api := app.Group("/api")

	api.Get("/sets", func(c *fiber.Ctx) error {
		sets := GetConfigContent()
		return c.JSON(sets)
	})
	
	app.Listen(":3000")
}