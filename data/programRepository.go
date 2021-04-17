package data

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func GetPrograms(db *pg.DB) ([]Program, error) {
	var programs []Program
	err := db.Model(&programs).
		Where("is_active = ?", true).
		Select()
	if err != nil {
		return nil, err
	}

	return programs, nil
}

func GetProgram(id int, db *pg.DB) (*Program, error) {
		program := &Program{}
		err := db.Model(program).
			Where("is_active = ?", true).
			Where("? = ?", pg.Ident("id"), id).
			Select()
		if err != nil {
			return &Program{}, err
		}
		return program, nil
}

func UpdateProgram(program *Program, db *pg.DB) (orm.Result, error) {
	res, err := db.Model(program).
		Where("? = ?", pg.Ident("id"), program.Id).
		UpdateNotZero()
	return res, err
}

func DeleteProgram(id int, db *pg.DB) (orm.Result, error) {
	program := Program{IsActive: false}
	res, err := db.Model(&program).
		Column("is_active").
		Where("? = ?", pg.Ident("id"), id).
		Update()
	return res, err
}