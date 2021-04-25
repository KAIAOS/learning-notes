package service

import (
	"fmt"
	"secKill/dao"
	"testing"
)

func TestReduceProductSendMsg(t *testing.T) {
	dao.InitMQchannel()
	err := ReduceProductSendMsg("hankai")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sssss")
}
