package mq

import (
	"log"
)

var done chan struct{}

// StartConsume 开始监听队列,获取消息
func StartConsume(qName, cName string, callback func(msg []byte) bool) {
	if channel == nil {
		log.Println("Initial channel failed.")
	}
	// 1. 通过 channel.Consume 获取消息信道
	msgs, err := channel.Consume(qName, cName, true, false, false, false, nil)
	if err != nil {
		log.Println(err)
		return
	}
	done = make(chan struct{})

	go func() {
		// 2. 循环获取队列的消息
		for msg := range msgs {
			// 3. 调用 callback 处理新的消息
			log.Println("Consumer has recive data", string(msg.Body), "from queue.")
			ok := callback(msg.Body)
			if !ok {
				// TODO: 写到另一个队列,用于异常情况的重试
			}
		}
	}()
	<-done

	channel.Close()
}
