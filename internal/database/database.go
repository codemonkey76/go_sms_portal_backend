package database

import (
	"database/sql"
	"fmt"
	"sms_portal/internal/env"
	"sync"

	_ "github.com/lib/pq"
)

var (
	dbOnce sync.Once
	db     *sql.DB
)

func InitDB() {
	var err error
	db, err = ConnectDB()

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
}

func GetDB() *sql.DB {
	fmt.Println("Getting DB")
	dbOnce.Do(InitDB)
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func ConnectDB() (*sql.DB, error) {
	connstr := env.GetConnectionString()
	fmt.Println("Connecting to DB: ", connstr)
	return sql.Open("postgres", env.GetConnectionString())
}
