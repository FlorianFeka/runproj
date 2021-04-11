package controllers

import (
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)


func RegisterControllers(api fiber.Router, db *pg.DB) {
	GetSets(api, db)
	GetSet(api, db)
	UpdateSet(api, db)
}