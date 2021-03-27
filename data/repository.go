package data

import (
	"github.com/go-pg/pg/v10"
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
