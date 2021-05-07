package controllers

import (
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterControllers(app fiber.Router, db *pg.DB) {
	RegisterSetControllers(app, db)
	RegisterProgramControllers(app, db)
	RegisterProgramSetControllers(app, db)
	RegisterArgumentControllers(app, db)
	RegisterSysteControllers(app, db);
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

func RegisterProgramSetControllers(app fiber.Router, db *pg.DB){
	GetProgramSets(app, db)
	GetProgramSet(app, db)
	UpdateProgramSet(app, db)
	CreateProgramSet(app, db)
	DeleteProgramSet(app, db)
}

func RegisterArgumentControllers(app fiber.Router, db *pg.DB){
	GetArguments(app, db)
	GetArgument(app, db)
	UpdateArgument(app, db)
	CreateArgument(app, db)
	DeleteArgument(app, db)
}

func RegisterSysteControllers(app fiber.Router, db *pg.DB) {
	ExecuteSet(app, db);
}