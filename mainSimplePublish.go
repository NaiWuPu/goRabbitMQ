package main

import (
	"fmt"
	"rabbitMQ/RabbitMQ"
)

func main(){
	rabbitmq := RabbitMQ.MewRabbitMQSimole("imoocSimple")
	rabbitmq.PublishSimple("Hello imooc!")
	 fmt.Printf("发送成功")
}
