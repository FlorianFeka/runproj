package controllers

import (
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)


func RegisterSetControllers(app *fiber.App, db *pg.DB) {
	api := app.Group("/api")
	GetSets(api, db)
	GetSet(api, db)
	UpdateSet(api, db)
	CreateSet(api, db)
	DeleteSet(api, db)
}