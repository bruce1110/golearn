package run

import "bruce.net/Servers/model"

func Init() {
	if err := model.InitMysql(); err != nil {
		panic("init mysql err" + err.Error())
	}
}
