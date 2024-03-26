package core

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sms_portal/env"
)

func ConnectDB() (*sql.DB, error) {
	return sql.Open("postgres", env.GetConnectionString())
}
