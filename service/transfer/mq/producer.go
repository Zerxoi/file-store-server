package mq

import (
	"file-store-server/config"
	"log"

	"github.com/streadway/amqp"
)

var coon *amqp.Connection
var channel *amqp.Channel

func init() {
	// 1. 获得rabbitmq的连接
	conn, err := amqp.Dial(config.RabbitURL)
	if err != nil {
		log.Println("initChannel:", err)
		return
	}

	// 2. 打开一个channel. 用于消息的发布与接受等
	channel, err = conn.Channel()
	if err != nil {
		log.Println("initChannel:", err)
		return
	}

}

// Publish 发布消息
func Publish(exchange, routingKey string, msg []byte) bool {
	// 1. 检查channel是否正常
	if channel == nil {
		log.Println("Initial channel failed.")
		return false
	}

	// 2. 执行消息发布动作
	err := channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	})
	if err != nil {
		log.Println("Publish:", err)
		return false
	}
	return true
}
