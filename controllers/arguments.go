package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FlorianFeka/runproj/data"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

func GetArguments(api fiber.Router, db *pg.DB) {
	api.Get("/arguments", func(c *fiber.Ctx) error {
		arguments, err := data.GetArguments(db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}
		return c.JSON(arguments)
	})
}

func GetArgument(api fiber.Router, db *pg.DB) {
	api.Get("/arguments/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		argument, err := data.GetArgument(id, db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}

		return c.JSON(*argument)
	})
}

func UpdateArgument(api fiber.Router, db *pg.DB) {
	api.Put("/arguments/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		argument := data.NewArgument("", 0, 0)
		err = json.Unmarshal(c.Request().Body(), &argument)
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}
		
		if id != argument.Id {
			c.Response().SetStatusCode(http.StatusBadRequest)
			if argument.Id == 0 {
				return c.JSON(data.ErrorResponse{
					FailedField: "Argument.Id",
					Tag: "required",
					Value: "",
				})
			}
			return err
		}
		
		_, err = data.UpdateArgument(&argument, db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		return nil
	})
}

func CreateArgument(api fiber.Router, db *pg.DB) {
	api.Post("/arguments", func(c *fiber.Ctx) error {
		argument := data.NewArgument("", 0, 0)
		err := json.Unmarshal(c.Request().Body(), &argument)
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		errors := data.ValidateArgument(argument)
		if errors != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return c.JSON(errors)
		}

		_, err = db.Model(&argument).Insert()
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}
		
		c.Response().SetStatusCode(http.StatusNoContent)

		return nil
	})
}

func DeleteArgument(api fiber.Router, db *pg.DB) {
	api.Delete("/arguments/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return err
		}

		_, err = data.DeleteArgument(id, db)
		if err != nil {
			c.Response().SetStatusCode(http.StatusInternalServerError)
			return err
		}

		c.Response().SetStatusCode(http.StatusNoContent)

		return nil
	})
}