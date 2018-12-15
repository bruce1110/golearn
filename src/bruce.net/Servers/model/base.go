package model

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var Db *sql.DB
var Pool *redis.Pool

const myconfig = "dev:123456@tcp(10.0.0.2:3306)/messenge"
const reconfig = "127.0.0.1:6379"

func InitMysql() error {
	var err error
	Db, err = sql.Open("mysql", myconfig)
	if err != nil {
		return err
	}
	Db.SetMaxIdleConns(1)
	Db.SetMaxOpenConns(2)
	Db.SetConnMaxLifetime(time.Second * 500)
	if err = Db.Ping(); err != nil {
		return err
	}
	return err
}

func InitRedis() error {
	var err error
	Pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", reconfig)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		MaxIdle:   1,
		MaxActive: 2,
	}
	return err
}
