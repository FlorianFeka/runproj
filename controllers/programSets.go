package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FlorianFeka/runproj/data"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

func GetProgramSets(api fiber.Router, db *pg.DB) {
	api.Get("/programSets", func(c *fiber.Ctx) error {
		programSets, err := data.GetProgramSets(db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}
		return c.JSON(programSets)
	})
}

func GetProgramSet(api fiber.Router, db *pg.DB) {
	api.Get("/programSets/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		programSet, err := data.GetProgramSet(id, db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}

		return c.JSON(*programSet)
	})
}

func UpdateProgramSet(api fiber.Router, db *pg.DB) {
	api.Put("/programSets/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		programSet := data.NewProgramSet(0, 0, 0, "")
		err = json.Unmarshal(c.Request().Body(), &programSet)
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}
		if id != programSet.Id {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		_, err = data.UpdateProgramSet(&programSet, db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		return nil
	})
}

func CreateProgramSet(api fiber.Router, db *pg.DB) {
	api.Post("/programSets", func(c *fiber.Ctx) error {
		programSet := data.NewProgramSet(0, 0, 0, "")
		err := json.Unmarshal(c.Request().Body(), &programSet)
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		_, err = db.Model(&programSet).Insert()
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}

		c.Response().SetStatusCode(http.StatusNoContent)

		return nil
	})
}

func DeleteProgramSet(api fiber.Router, db *pg.DB) {
	api.Delete("/programSets/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		_, err = data.DeleteProgramSet(id, db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}

		c.Response().SetStatusCode(http.StatusNoContent)

		return nil
	})
}
