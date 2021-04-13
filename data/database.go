package data

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	// init postgres driver
	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	_ "github.com/lib/pq"
)

var (
	host            = "localhost"
	port            = 5432
	user            = "runproj"
	password        = "mysecretpassword"
	dbname          = "runproj"
	maxConnectTries = 3
	tablesExisted   = false
	debug           = true
)

func GetPgDbConnection() *pg.DB {
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
	if debug {
		db.AddQueryHook(pgdebug.DebugHook{
			Verbose: true,
		})
	}

	return db
}

func SetConfFromEnv() {
	debugModeEnv, present := os.LookupEnv("DEBUG")
	if present {
		debugMode, err := strconv.ParseBool(debugModeEnv)
		if err != nil {
			panic(err)
		}
		debug = debugMode
	}
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
	if present {
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

	if !tablesExisted && debug {
		PopulateTestData(db)
	}
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
			tablesExisted = true
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
		Name:     "timeo flutter",
		IsActive: true,
	}

	timeoAngular := Set{
		Name:     "timeo angular",
		IsActive: true,
	}

	_, err := db.Model(&[]*Set{&timeoAngular, &timeoFlutter}).Insert()
	if err != nil {
		panic(err)
	}

	vsCode := Program{
		Name:        "Visual Studio Code",
		ProgramPath: "C:\\Users\\Feka\\AppData\\Local\\Programs\\Microsoft VS Code\\Code.exe",
		IsActive:    true,
	}

	chrome := Program{
		Name:        "Chrome",
		ProgramPath: "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
		IsActive:    true,
	}

	_, err = db.Model(&[]*Program{&vsCode, &chrome}).Insert()
	if err != nil {
		panic(err)
	}

	timeoFlutterVsCode := ProgramSet{
		SetId:     timeoFlutter.Id,
		ProgramId: vsCode.Id,
		IsActive:  true,
	}

	timeoFlutterChrome := ProgramSet{
		SetId:     timeoFlutter.Id,
		ProgramId: chrome.Id,
		IsActive:  true,
	}

	timeoAngularVsCode := ProgramSet{
		SetId:     timeoAngular.Id,
		ProgramId: vsCode.Id,
		IsActive:  true,
	}

	timeoAngularChrome := ProgramSet{
		SetId:     timeoAngular.Id,
		ProgramId: chrome.Id,
		IsActive:  true,
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
			IsActive:     true,
		},
		{
			Argument:     "--new-window",
			ProgramSetId: timeoFlutterChrome.Id,
			IsActive:     true,
		},
		{
			Argument:     "https://github.com/FlorianFeka/timeo-app",
			ProgramSetId: timeoFlutterChrome.Id,
			IsActive:     true,
		},
		{
			Argument:     "D:\\2_Projekte\\Javascript\\Angular\\timeo",
			ProgramSetId: timeoAngularVsCode.Id,
			IsActive:     true,
		},
		{
			Argument:     "--new-window",
			ProgramSetId: timeoAngularVsCode.Id,
			IsActive:     true,
		},
		{
			Argument:     "https://github.com/FlorianFeka/timeo",
			ProgramSetId: timeoAngularVsCode.Id,
			IsActive:     true,
		}}

	_, err = db.Model(&arr).Insert()
	if err != nil {
		panic(err)
	}
}
