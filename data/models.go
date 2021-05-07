package data

import "github.com/go-playground/validator"

type Set struct {
	Id          int           `pg:",pk"`
	Name        string        `validate:"required"`
	ProgramSets []*ProgramSet `pg:"rel:has-many" json:",omitempty"`
	IsActive    bool          `json:"-"`
}

func NewSet(name string) Set {
	return Set{
		Name:     name,
		IsActive: true,
	}
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateSet(set Set) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	if err := validate.Struct(set); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

type Program struct {
	Id          int    `pg:",pk"`
	Name        string `validate:"required"`
	ProgramPath string `validate:"required"`
	IsActive    bool   `json:"-"`
}

func NewProgram(name, programPath string) Program {
	return Program{
		Name:        name,
		ProgramPath: programPath,
		IsActive:    true,
	}
}

func ValidateProgram(program Program) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	if err := validate.Struct(program); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

type Argument struct {
	Id           int         `pg:",pk"`
	Argument     string      `validate:"required"`
	Order        int         `validate:"required"`
	ProgramSetId int         `validate:"required"`
	ProgramSet   *ProgramSet `pg:"rel:has-one" json:"-"`
	IsActive     bool        `json:"-"`
}

func NewArgument(
	argument string,
	order, programSetId int) Argument {
	return Argument{
		Argument:     argument,
		Order:        order,
		ProgramSetId: programSetId,
		IsActive:     true,
	}
}

func ValidateArgument(argument Argument) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	if err := validate.Struct(argument); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

type ProgramSet struct {
	Id              int      `pg:",pk"`
	SetId           int      `validate:"required"`
	Set             *Set     `pg:"rel:has-one"`
	ProgramId       int      `validate:"required"`
	Program         *Program `pg:"rel:has-one"`
	Monitor         int
	SnappedPosition string
	Arguments       []*Argument `pg:"rel:has-many"`
	IsActive        bool        `json:"-"`
}

func NewProgramSet(
	setId, programId, monitor int,
	snappedPosition string) ProgramSet {
	return ProgramSet{
		SetId:           setId,
		ProgramId:       programId,
		Monitor:         monitor,
		SnappedPosition: snappedPosition,
		IsActive:        true,
	}
}

func ValidateProgramSet(programSet ProgramSet) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	if err := validate.Struct(programSet); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
