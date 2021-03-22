package data

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "mysecretpassword"
  dbname   = "runproj_test"
)

func CreateDatabase(){
	db := pg.Connect(&pg.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
		User: user,
		Password: password,
		Database: dbname,
	})
	defer db.Close()
	CreateSchema(db)
}

func CreateSchema(db *pg.DB) error {
	models := []interface{} {
		(*Set)(nil),
		(*Program)(nil),
		(*Argument)(nil),
		(*ProgramSet)(nil),
	};

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}