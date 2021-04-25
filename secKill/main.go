package main

import (
	"secKill/dao"
	"secKill/router"
	"secKill/service"
)

func main(){
	dao.InitMysql()
	dao.InitRedis()
	dao.InitMQchannel()
	defer dao.Close()

	//开启消费者线程
	go service.StartConsume()

	r := router.SetUpRouter()

	r.Run(":8000")
}