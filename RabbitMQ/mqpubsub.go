package RabbitMQ

import (
	"github.com/streadway/amqp"
	"log"
)

// 订阅模式创建RabbitMQ实例
func NewRabbitMQPubSub(queueName, exchangeName, key string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, exchangeName, key)
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")

	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel!")
	return rabbitmq
}

// 订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
	// 1. 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an exchange")
	// 2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// 订阅模式消费代码
func (r *RabbitMQ) RecieveSub() {
	// 1. 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an exchange")
	// 2 试探性创建队列，队列名字不要写，随机拿到名字
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")
	// 绑定队列到 exchange 中
	err = r.channel.QueueBind(
		// 获取上面随机生成的名字
		q.Name,
		// 再pub.sub 模式下，key 为空
		r.Key,
		r.Exchange,
		false,
		nil,
	)

	// 消费消息
	messges, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
		go func() {
			for d := range messges {
				log.Printf("Received a message: %s", d.Body)
			}
		}()
	//fmt.Println("退出")
	<-forever
}
