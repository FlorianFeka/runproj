package data

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func GetArguments(db *pg.DB) ([]Argument, error) {
	var programs []Argument
	err := db.Model(&programs).
		Where("is_active = ?", true).
		Select()
	if err != nil {
		return nil, err
	}

	return programs, nil
}

func GetArgument(id int, db *pg.DB) (*Argument, error) {
		program := &Argument{}
		err := db.Model(program).
			Where("is_active = ?", true).
			Where("? = ?", pg.Ident("id"), id).
			Select()
		if err != nil {
			return &Argument{}, err
		}
		return program, nil
}

func UpdateArgument(program *Argument, db *pg.DB) (orm.Result, error) {
	res, err := db.Model(program).
		Where("? = ?", pg.Ident("id"), program.Id).
		UpdateNotZero()
	return res, err
}

func DeleteArgument(id int, db *pg.DB) (orm.Result, error) {
	res, err := db.Model(&Argument{}).
		Set("is_active = ?", false).
		Where("? = ?", pg.Ident("id"), id).
		Update()
	return res, err
}