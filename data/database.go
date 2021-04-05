package data

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	// init postgres driver
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "runproj"
	password = "mysecretpassword"
	dbname   = "runproj"
	maxConnectTries = 3
)

func GetPgDbConnection() (*pg.DB) {
	SetConfFromEnv()
	tries := 1
	connect: 
	opt, err := pg.ParseURL(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname))
	if err != nil {
		panic(err)
	}
	
	ctx := context.Background()

	db := pg.Connect(opt)

	if err := db.Ping(ctx); err != nil {
		if tries >= maxConnectTries {
			panic(fmt.Sprintf(
				"Couldn't connect to db '%s' after %d tries\n\n\n%s", 
				dbname, 
				tries, 
				err.Error()))
		}
		fmt.Printf("Couldn't connect to db '%s'. Current try: %d\n", dbname, tries)
		time.Sleep(5 * time.Second)
		tries++
		goto connect
	}

	return db
}

func SetConfFromEnv() {
	dbHost, present := os.LookupEnv("DB_HOST")
	if present {
		host = dbHost
	}
	dbUser, present := os.LookupEnv("DB_USER")
	if present {
		user = dbUser
	}
	dbPassword, present := os.LookupEnv("DB_PASSWORD")
	if present {
		password = dbPassword
	}
	dbDatabase, present := os.LookupEnv("DB_DATABASE")
	if present {
		dbname = dbDatabase
	}
	dbPort, present := os.LookupEnv("DB_PORT")
	if present {
		dbIntPort, err := strconv.Atoi(dbPort)
		if err != nil {
			panic(err)
		}
		port = dbIntPort
	}
	maxTries, present := os.LookupEnv("DB_CONNECT_MAX_TRIES")
	if present{
		dbIntMaxTries, err := strconv.Atoi(maxTries)
		if err != nil {
			panic(err)
		}
		maxConnectTries = dbIntMaxTries
	}
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
