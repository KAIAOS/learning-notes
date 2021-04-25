package service

import (
	"secKill/dao"
	"testing"
)

func TestStartConsume(t *testing.T) {
	dao.InitMysql()
	dao.InitMQchannel()
	StartConsume()
}
