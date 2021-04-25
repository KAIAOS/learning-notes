package service

import (
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"log"
	"secKill/dao"
)
type MsgOrder struct {
	ProductId int `json:"product_id"`
	UserName string `json:"user_name"`
}

func ReduceProductSendMsg(userName string)(err error){

	result, err := dao.Redisdb.Decr("product").Result()
	if result <0 {
		log.Printf("卖完了 result:%d", result)
		err = errors.New("卖完了")
		return
	}

	var data = MsgOrder{
		ProductId: int(result),
		UserName: userName,
	}

	body, _ := json.Marshal(data)

	err = dao.MQCH.Publish(
		"",
		dao.QNAME,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/palin",
			Body: body,
		})
	if err != nil{
		log.Fatalf("Failed to publish a message: %s", err)
		return
	}
	return
}