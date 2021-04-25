package dao

import (
	"github.com/streadway/amqp"
	"log"
)

var (
	conn  *amqp.Connection
	MQCH  *amqp.Channel
	QNAME string
)

func InitMQchannel()(err error){
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil{
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		return
	}
	MQCH, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
		return
	}
	q, err := MQCH.QueueDeclare(
		"orders", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
		return
	}
	QNAME = q.Name
	return
}