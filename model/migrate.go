package db

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	h "github.com/taufikardiyan28/mataharitest1/helper"
)

type CekTable struct {
	Count int `db:"cnt"`
}

func Migrate() error {
	dbConfig := mysql.Config{
		User:                 h.Config.Database.User,
		Passwd:               h.Config.Database.Password,
		Loc:                  time.Local,
		Net:                  fmt.Sprintf("tcp(%s:%d)", h.Config.Database.Host, h.Config.Database.Port),
		AllowNativePasswords: true,
	}

	con, err := sqlx.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		return err
	}

	dbName := h.Config.Database.DbName
	strSQL := fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s;`, dbName)
	if _, err = con.Exec(strSQL); err != nil {
		return err
	}

	if _, err = con.Exec(fmt.Sprintf(`USE %s`, dbName)); err != nil {
		return err
	}

	if err := CreateProducts(con, dbName); err != nil {
		return err
	}

	err = CreateCarts(con, dbName)

	con.Close()
	return err
}

func CreateProducts(con *sqlx.DB, dbName string) error {
	t := CekTable{}

	strSQL := `SELECT COUNT(*) cnt  
				FROM information_schema.tables
				WHERE table_schema = ? 
					AND table_name = 'products'
				LIMIT 1;`
	err := con.Get(&t, strSQL, dbName)
	if err != nil {
		return err
	}
	if t.Count > 0 {
		return nil
	}

	strSQL = `CREATE TABLE products (
		id int(11) NOT NULL AUTO_INCREMENT,
		productName varchar(100) NOT NULL,
		stock int(11) NOT NULL DEFAULT '0',
		createdAt datetime DEFAULT CURRENT_TIMESTAMP,
		updatedAt datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	  ) ENGINE=InnoDB`

	if _, err = con.Exec(strSQL); err != nil {
		return err
	}

	strSQL = `insert  into products(productName, stock) values 
				('Topi X',10),
				('Topi Y',10),
				('Topi Z',10),
				('T-Shirt X',5),
				('T-Shirt Y',5),
				('T-Shirt Z',5);`

	_, err = con.Exec(strSQL)
	return err
}

func CreateCarts(db *sqlx.DB, dbName string) error {
	t := CekTable{}

	strSQL := `SELECT COUNT(*) cnt  
				FROM information_schema.tables
				WHERE table_schema = ? 
					AND table_name = 'carts'
				LIMIT 1;`
	err := db.Get(&t, strSQL, dbName)
	if err != nil {
		return err
	}
	if t.Count > 0 {
		return nil
	}

	strSQL = `CREATE TABLE carts (
		id int(11) NOT NULL AUTO_INCREMENT,
		productId int(11) NOT NULL,
		qty int(11) NOT NULL,
		AID char(32) DEFAULT '',
		userId int(11) DEFAULT NULL,
		isOrdered int(1) NOT NULL DEFAULT '0',
		createdAt datetime DEFAULT CURRENT_TIMESTAMP,
		updatedAt datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		KEY productId (productId),
		CONSTRAINT carts_ibfk_1 FOREIGN KEY (productId) REFERENCES products (id) ON UPDATE CASCADE
	  ) ENGINE=InnoDB`

	_, err = db.Exec(strSQL)
	return err
}
