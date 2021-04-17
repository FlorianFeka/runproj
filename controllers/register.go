package controllers

import (
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterControllers(app fiber.Router, db *pg.DB) {
	RegisterSetControllers(app, db)
	RegisterProgramControllers(app, db)
}

func RegisterSetControllers(app fiber.Router, db *pg.DB) {
	GetSets(app, db)
	GetSet(app, db)
	UpdateSet(app, db)
	CreateSet(app, db)
	DeleteSet(app, db)
}

func RegisterProgramControllers(app fiber.Router, db *pg.DB) {
	GetPrograms(app, db)
	GetProgram(app, db)
	UpdateProgram(app, db)
	CreateProgram(app, db)
	DeleteProgram(app, db)
}