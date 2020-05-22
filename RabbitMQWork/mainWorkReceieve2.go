package main

import "rabbitMQ/RabbitMQ"

func main(){
	rabbitmq := RabbitMQ.MewRabbitMQSimole("imoocSimple")
	rabbitmq.ConsumeSimple()
}