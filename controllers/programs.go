package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/FlorianFeka/runproj/data"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

func GetPrograms(api fiber.Router, db *pg.DB) {
	api.Get("/programs", func(c *fiber.Ctx) error {
		programs, err := data.GetPrograms(db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}
		return c.JSON(programs)
	})
}

func GetProgram(api fiber.Router, db *pg.DB) {
	api.Get("/programs/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		program, err := data.GetProgram(id, db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}

		return c.JSON(*program)
	})
}

func UpdateProgram(api fiber.Router, db *pg.DB) {
	api.Put("/programs/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		program := data.NewProgram("", "")
		err = json.Unmarshal(c.Request().Body(), &program)
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}
		
		if id != program.Id {
			c.Response().SetStatusCode(http.StatusBadRequest)
			if program.Id == 0 {
				return c.JSON(data.ErrorResponse{
					FailedField: "Program.Id",
					Tag: "required",
					Value: "",
				})
			}
			return err
		}
		
		_, err = data.UpdateProgram(&program, db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		return nil
	})
}

func CreateProgram(api fiber.Router, db *pg.DB) {
	api.Post("/programs", func(c *fiber.Ctx) error {
		program := data.NewProgram("", "")
		err := json.Unmarshal(c.Request().Body(), &program)
		if err != nil ||  strings.TrimSpace(program.Name) == "" {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		errors := data.ValidateProgram(program)
		if errors != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return c.JSON(errors)
		}

		_, err = db.Model(&program).Insert()
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}
		
		c.Response().SetStatusCode(http.StatusNoContent)

		return nil
	})
}

func DeleteProgram(api fiber.Router, db *pg.DB) {
	api.Delete("/programs/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		_, err = data.DeleteProgram(id, db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}

		c.Response().SetStatusCode(http.StatusNoContent)

		return nil
	})
}