package database

import (
	"database/sql"
	"sms_portal/core"
	"sync"
)

var (
	dbOnce sync.Once
	db     *sql.DB
)

func InitDB() {
	var err error
	db, err = core.ConnectDB()

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
