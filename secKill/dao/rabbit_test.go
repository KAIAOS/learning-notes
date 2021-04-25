package dao

import (
	"fmt"
	"testing"
)

func TestInitMQchannel(t *testing.T) {
	err := InitMQchannel()
	if err != nil {
		panic(err)
	}
	fmt.Println("success!")
}
