package data

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func GetProgramSets(db *pg.DB) ([]ProgramSet, error) {
	var programSets []ProgramSet
	err := db.Model(&programSets).
		Relation("Set").
		Relation("Program").
		Relation("Arguments").
		Where("program_set.is_active = ?", true).
		Select()
	if err != nil {
		return nil, err
	}

	return programSets, nil
}

func GetProgramSet(id int, db *pg.DB) (*ProgramSet, error) {
	programSet := &ProgramSet{}
	err := db.Model(programSet).
		Relation("Set").
		Relation("Program").
		Relation("Arguments").
		Where("program_set.is_active = ?", true).
		Where("? = ?", pg.Ident("program_set.id"), id).
		Select()
	if err != nil {
		return &ProgramSet{}, err
	}
	return programSet, nil
}

func UpdateProgramSet(programSet *ProgramSet, db *pg.DB) (orm.Result, error) {
	res, err := db.Model(programSet).
		Where("? = ?", pg.Ident("id"), programSet.Id).
		UpdateNotZero()
	return res, err
}

func DeleteProgramSet(id int, db *pg.DB) (orm.Result, error) {
	res, err := db.Model(&ProgramSet{}).
		Set("is_active = ?", false).
		Where("? = ?", pg.Ident("id"), id).
		Update()
	return res, err
}
