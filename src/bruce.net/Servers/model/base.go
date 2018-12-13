package model

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var Db *sql.DB

const config = "pepper_passport:0b570fc5108d7bf1@tcp(10.138.67.99:2146)/pepper_passport"

func InitMysql() error {
	var err error
	Db, err = sql.Open("mysql", config)
	if err != nil {
		return err
	}
	Db.SetMaxIdleConns(1)
	Db.SetMaxOpenConns(2)
	if err = Db.Ping(); err != nil {
		return err
	}
	return err
}
