package main

import (
	"github.com/astaxie/beego/logs"
	"rabbitMQ/RabbitMQ"
	"strconv"
)

func main() {
	rabbitmq:= RabbitMQ.NewRabbitMQPubSub("wa","chijiuhua", "")
	for i := 0; i < 10000; i++ {
		rabbitmq.PublishPub(strconv.Itoa(i))
		//time.Sleep(1 * time.Second)
		logs.Debug(i)
	}
}