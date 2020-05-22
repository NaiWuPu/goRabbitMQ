package main

import "rabbitMQ/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting(
		"member", "member_login")
	imoocOne.RecieveRouting()
}