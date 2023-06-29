package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	cfg := mysql.Config{
		User:                 "sample_database",
		Passwd:               "sample",
		Net:                  "tcp",
		Addr:                 "mysql:3306",
		DBName:               "sample_database",
		AllowNativePasswords: true,
	}

	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to mysql...")
}
