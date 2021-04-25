package dao

import (
	"fmt"
	"secKill/model"
	"testing"
)

func init(){
	InitMysql()
}

func TestInsertOrders(t *testing.T) {
	var order = &model.Orders{
		ProductId: 2,
		UserName: "hankai",
	}
	orders, err := InsertOrders(order)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(orders)
}

func TestGetOrdersById(t *testing.T) {
	order, err := GetOrdersById(3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", len(order))
}