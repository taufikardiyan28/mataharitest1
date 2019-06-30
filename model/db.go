package db

import (
	"github.com/jmoiron/sqlx"
)

type DB struct{}

var Pool *sqlx.DB

func InitDB(dsn string) error {
	var err error
	Pool, err = sqlx.Open("mysql", dsn)

	return err
}

func (d *DB) GetPool() *sqlx.DB {
	return Pool
}
