package main

import "rabbitMQ/RabbitMQ"

func main()  {
	rabbitmq:= RabbitMQ.NewRabbitMQPubSub("wa","chijiuhua", "")
	rabbitmq.RecieveSub()
}
