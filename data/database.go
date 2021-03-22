package data

import (
	"database/sql"
	"fmt"

	// init postgres driver
	_ "github.com/lib/pq"
)

// TestDB connects to database and outputs if successful connection was created
func TestDB() {	
  	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s  sslmode=disable",
    host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

// CreateDB deletes DB if exists and creates DB 
func CreateDB(){
  	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s sslmode=disable",
    host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;\n"+
		"CREATE DATABASE %s;", dbname, dbname))
	if err != nil {
		panic(err)
	}

}