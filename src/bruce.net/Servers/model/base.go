package model

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/garyburd/redigo/redis"

var Db *sql.DB
var Pool *redis.Pool

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

func InitRedis() error {
	Pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		MaxIdle:   1,
		MaxActive: 2,
	}
	return nil
}
