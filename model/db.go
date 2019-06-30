package db

import (
	"github.com/jmoiron/sqlx"
)

var Pool *sqlx.DB

func InitDB(dsn string) error {
	var err error
	Pool, err = sqlx.Open("mysql", dsn)

	return err
}
