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
		User:                 "dr_stone",
		Passwd:               "senku",
		Net:                  "tcp",
		Addr:                 "mysql:3306",
		DBName:               "dr_stone",
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

func GetDB() *sql.DB {
	return db
}