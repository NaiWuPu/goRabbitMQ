package main

import "rabbitMQ/RabbitMQ"

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	rabbitmq.RecieveSub()
}
