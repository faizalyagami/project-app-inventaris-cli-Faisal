package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb()  {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=inventaris_kantor sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Koneksi Database berhasil.")
	}

}

func GetDb() *sql.DB {
	return db
}