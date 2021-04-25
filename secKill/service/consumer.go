package service

import (
	"encoding/json"
	"log"
	"secKill/dao"
	"secKill/model"
)

func StartConsume(){
	log.Print("enter consumer")
	msgs, err := dao.MQCH.Consume(
		dao.QNAME, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}
	forever := make(chan bool)

	go func(){
		for d := range msgs{
			var data MsgOrder
			_ = json.Unmarshal(d.Body, &data)
			//fmt.Printf("%v\n",data)
			orders, _ := dao.GetOrdersById(data.ProductId)
			if len(orders) == 0{
				var o = &model.Orders{
					ProductId: data.ProductId,
					UserName: data.UserName,
				}

				_, err = dao.InsertOrders(o)
				if err != nil{
					log.Printf("failed to inser into orders: %s",err)
				}
				log.Printf("got one product user: %s, product id: %d\n", data.UserName,data.ProductId)
			}
		}
	}()

	log.Printf("[*] waiting for messages.")
	<-forever
}



