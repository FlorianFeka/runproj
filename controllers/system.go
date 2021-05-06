package controllers

import (
	"net/http"
	"strconv"

	"github.com/FlorianFeka/runproj/actions"
	"github.com/FlorianFeka/runproj/data"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

func ExecuteSet(api fiber.Router, db *pg.DB) {
	api.Post("/executeSet/:setId", func(c *fiber.Ctx) error {
		setId, err := strconv.Atoi(c.Params("setId"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		set, err := data.GetFullSet(setId, db)

		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		actions.ExecuteSet(set)

		return c.JSON(set)
	})
}
