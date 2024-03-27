package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sms_portal/internal/env"
	"sync"
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
	dbOnce.Do(InitDB)
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func ConnectDB() (*sql.DB, error) {
	return sql.Open("postgres", env.GetConnectionString())
}
