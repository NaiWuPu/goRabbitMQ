package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)


// 简单模式 Step1: 创建简单的MQ
func MewRabbitMQSimole(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

func (r *RabbitMQ) PublishSimple(message string) {
	// 1. 申请队列,如果队列不存在会自动创建,如果存在则创建
	// 保证队列存在，详细发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 消息是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		fmt.Print(err)
	}

	// 2. 发送消息到队列中
	_ = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// 如果为 true 根据exhange类型和routkey规则，如果无法找到符合条件的队列会把发送的消息返回给发送者
		false,
		// 如果为true, 当exchange 发送消息到队列后发现队列上没有绑定消费者，则会把消息发挥给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

}

// 消费端
func (r *RabbitMQ) ConsumeSimple() {
	// 1. 申请队列,如果队列不存在会自动创建,如果存在则创建
	// 保证队列存在，详细发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 消息是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		fmt.Print(err)
	}

	// 接受消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		// 用来区分多个消费者
		"",
		// 是否自动应答
		true,
		// 是否自动应答
		false,
		// 如果设置为true, 表示不能将同一个connection 中发送的消息传递给这个connection 中的消费者
		false,
		// 消费是否为阻塞
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)

	// 启用携程处理消息
	go func() {
		for d := range msgs {
			// 实现处理逻辑
			log.Printf("Received a message :%s", d.Body)
		}
	}()

	<-forever
}

