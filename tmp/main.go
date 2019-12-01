package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tietang/dbx"
	"time"
)



func main() {
	settings := dbx.Settings{
		DriverName:      "mysql",
		User:            "root",
		Password:        "12345678",
		Host:            "127.0.0.1:3306",
		MaxOpenConns:    10,
		MaxIdleConns:    2,
		ConnMaxLifetime: time.Minute * 30,
		Options: map[string]string{
			"charset":   "utf8",
			"parseTime": "true",
		},
	}
	db, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}

}