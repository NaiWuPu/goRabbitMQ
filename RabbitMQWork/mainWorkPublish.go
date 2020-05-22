package main

import (
	"fmt"
	"rabbitMQ/RabbitMQ"
	"strconv"
	"time"
)

func main(){
	rabbitmq := RabbitMQ.MewRabbitMQSimole("imoocSimple")
	for i := 0; i <= 100 ; i++ {
		rabbitmq.PublishSimple("Hello imooc!"+ strconv.Itoa(i))
		time.Sleep(1 *time.Second)
		fmt.Println( i)
	}
}
