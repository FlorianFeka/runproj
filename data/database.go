package data

import (
	"database/sql"
	"fmt"

	// init postgres driver
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "runproj_test"
)

func GetPgDbConnection() (*pg.DB) {
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		User:     user,
		Password: password,
		Database: dbname,
	})
	return db
}

func CreateDatabase(db *pg.DB) {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s  sslmode=disable",
	// 	host, port, user, password)
	// rawDb, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }

	// defer rawDb.Close()

	// DeleteExistingDatabase(rawDb)
	CreateSchema(db)
	PopulateTestData(db)
}

func CreateSchema(db *pg.DB) {
	models := []interface{}{
		(*Set)(nil),
		(*Program)(nil),
		(*ProgramSet)(nil),
		(*Argument)(nil),
	}

	for _, model := range models {
		exists, err := db.Model(model).Exists()

		if err != nil {
			fmt.Println(err)
		}

		if exists {
			continue
		}

		modelName := fmt.Sprintf("%T", model)
		fmt.Println("Create table:", modelName)

		err = db.Model(model).CreateTable(&orm.CreateTableOptions{
			FKConstraints: true,
		})
		if err != nil {
			panic(err)
		}
	}
}

func DeleteExistingDatabase(db *sql.DB) {
	_, err := db.Exec(`drop database %s WITH (FORCE);
						create database %s;`, dbname, dbname)
	if err != nil {
		panic(err)
	}
}

func PopulateTestData(db *pg.DB) {
	timeoFlutter := Set{
		Name: "timeo flutter",
	}

	timeoAngular := Set{
		Name: "timeo angular",
	}

	_, err := db.Model(&[]*Set{&timeoAngular, &timeoFlutter}).Insert()
	if err != nil {
		panic(err)
	}

	vsCode := Program{
		Name:        "Visual Studio Code",
		ProgramPath: "C:\\Users\\Feka\\AppData\\Local\\Programs\\Microsoft VS Code\\Code.exe",
	}

	chrome := Program{
		Name:        "Chrome",
		ProgramPath: "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
	}

	_, err = db.Model(&[]*Program{&vsCode, &chrome}).Insert()
	if err != nil {
		panic(err)
	}

	timeoFlutterVsCode := ProgramSet{
		SetId:     timeoFlutter.Id,
		ProgramId: vsCode.Id,
	}

	timeoFlutterChrome := ProgramSet{
		SetId:     timeoFlutter.Id,
		ProgramId: chrome.Id,
	}

	timeoAngularVsCode := ProgramSet{
		SetId:     timeoAngular.Id,
		ProgramId: vsCode.Id,
	}

	timeoAngularChrome := ProgramSet{
		SetId:     timeoAngular.Id,
		ProgramId: chrome.Id,
	}

	_, err = db.Model(&[]*ProgramSet{
		&timeoFlutterVsCode,
		&timeoFlutterChrome,
		&timeoAngularVsCode,
		&timeoAngularChrome,
		}).Insert()
	if err != nil {
		panic(err)
	}

	arr := []*Argument{
		{
			Argument:     "D:\\2_Projekte\\Flutter\\timeo",
			ProgramSetId: timeoFlutterVsCode.Id,
		},
		{
			Argument:     "--new-window",
			ProgramSetId: timeoFlutterChrome.Id,
		},
		{
			Argument:     "https://github.com/FlorianFeka/timeo-app",
			ProgramSetId: timeoFlutterChrome.Id,
		},
		{
			Argument: "D:\\2_Projekte\\Javascript\\Angular\\timeo",
			ProgramSetId: timeoAngularVsCode.Id,
		},
		{
			Argument: "--new-window",
			ProgramSetId: timeoAngularVsCode.Id,
		},
		{
			Argument: "https://github.com/FlorianFeka/timeo",
			ProgramSetId: timeoAngularVsCode.Id,
		}}

	_, err = db.Model(&arr).Insert()
	if err != nil {
		panic(err)
	}
}
