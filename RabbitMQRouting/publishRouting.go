package main

import (
	"fmt"
	"rabbitMQ/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting(
		"member", "member_login")
	for i := 0; i <= 1000; i++ {
		imoocOne.PublishRouting("immoc Routing one "+ strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
