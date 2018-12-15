package model

import (
	"github.com/garyburd/redigo/redis"
)

func GetOne(uid uint64) (string, string, error) {
	sql := "select nickname from user_1 where uid = ?";
	var name string
	Db.QueryRow(sql, uid).Scan(&name)
	rc := Pool.Get()
	defer rc.Close()
	value, _ := redis.String(rc.Do("GET", "test"))
	return name, value, nil
}
