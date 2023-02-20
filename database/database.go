package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Instance *sql.DB

const dbName = "astrofirechat"
const dbAddress = ""
const dbUser = "astrofirechat"
const dbPassword = "K1$qHRD0QTRpJv8ReaMW"
const dbParameters = "?parseTime=true"

func init() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s%s", dbUser, dbPassword, dbAddress, dbName, dbParameters))
	if err != nil {
		panic(err)
	}
	Instance = db
}
