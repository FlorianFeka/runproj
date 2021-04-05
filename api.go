package main

import (
	"fmt"

	"github.com/FlorianFeka/runproj/controllers"
	"github.com/FlorianFeka/runproj/data"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := data.GetPgDbConnection()
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	data.CreateDatabase(db)
	StartAPI(db)
	// args := os.Args[1:]

	// sets := GetConfigContent()

	// ExecuteSelectedSets(sets, args)
}

// StartAPI API for runproj
func StartAPI(db *pg.DB) {
	app := fiber.New()

	api := app.Group("/api")

	controllers.RegisterControllers(api, db)

	app.Listen(":3000")
}
