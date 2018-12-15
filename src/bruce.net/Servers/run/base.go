package run

import "bruce.net/Servers/model"

func Init() {
	if err := model.InitMysql(); err != nil {
		panic("init mysql err" + err.Error())
	}
	if err := model.InitRedis(); err != nil {
		panic("init redis err" + err.Error())
	}
}
