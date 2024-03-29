package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FlorianFeka/runproj/data"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

func GetSets(api fiber.Router, db *pg.DB) {
	api.Get("/sets", func(c *fiber.Ctx) error {
		sets, err := data.GetSets(db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}
		return c.JSON(sets)
	})
}

func GetSet(api fiber.Router, db *pg.DB) {
	api.Get("/sets/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		set, err := data.GetSet(id, db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}

		return c.JSON(*set)
	})
}

func UpdateSet(api fiber.Router, db *pg.DB) {
	api.Put("/sets/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		set := data.NewSet("")
		err = json.Unmarshal(c.Request().Body(), &set)
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		errors := data.ValidateSet(set)
		if errors != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return c.JSON(errors)
		}

		if id != set.Id {
			c.Response().SetStatusCode(http.StatusBadRequest)
			if set.Id == 0 {
				return c.JSON(data.ErrorResponse{
					FailedField: "Set.Id",
					Tag: "required",
					Value: "",
				})
			}
			return nil
		}
		
		_, err = data.UpdateSet(&set, db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		return nil
	})
}

func CreateSet(api fiber.Router, db *pg.DB) {
	api.Post("/sets", func(c *fiber.Ctx) error {
		set := data.NewSet("")
		err := json.Unmarshal(c.Request().Body(), &set)
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		errors := data.ValidateSet(set)
		if errors != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return c.JSON(errors)
		}
		
		_, err = db.Model(&set).Insert()
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}
		
		c.Response().SetStatusCode(http.StatusNoContent)

		return nil
	})
}

func DeleteSet(api fiber.Router, db *pg.DB) {
	api.Delete("/sets/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		_, err = data.DeleteSet(id, db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}

		c.Response().SetStatusCode(http.StatusNoContent)

		return nil
	})
}