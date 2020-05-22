package main

import (
	"rabbitMQ/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitmq:= RabbitMQ.NewRabbitMQPubSub("newProduct")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub("订阅模式生产 :"+strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}