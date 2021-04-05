package controllers

import (
	"strconv"

	"github.com/FlorianFeka/runproj/data"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

func GetSets(api fiber.Router, db *pg.DB) {
	api.Get("/sets", func(c *fiber.Ctx) error {
		sets, err := data.GetSets(db)
		if err != nil {
			return err
		}
		return c.JSON(sets)
	})
}

func GetSet(api fiber.Router, db *pg.DB) {
	api.Get("/sets/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(500)
			return err
		}
		set := data.Set{}
		err = db.Model(&set).
			Where("? = ?", pg.Ident("id"), id).
			Select()
		if err != nil {
			return err
		}

		return c.JSON(set)
	})
}