package data

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func GetSets(db *pg.DB) ([]Set, error) {
	var sets []Set
	err := db.Model(&sets).
		Select()
	if err != nil {
		return nil, err
	}

	return sets, nil
}

func GetSet(id int, db *pg.DB) (*Set, error) {
		set := &Set{}
		err := db.Model(set).
			Where("? = ?", pg.Ident("id"), id).
			Select()
		if err != nil {
			return &Set{}, err
		}
		return set, nil
}

func UpdateSet(set *Set, db *pg.DB) (orm.Result, error) {
	res, err := db.Model(set).
		Where("? = ?", pg.Ident("id"), set.Id).
		UpdateNotZero()
	return res, err
}
