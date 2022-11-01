package controllers

import (
	"CobraApp/pkg/rabbitmq"
	"fmt"
	"log"
	"strconv"
	"time"
)

// SendMQ 写入队列
func SendMQ() {
	rabbitmqRes := rabbitmq.NewRabbitMQSimple("golangSimple")

	for i := 0; i < 10; i++ {
		rabbitmqRes.PublishSimple("Hello golang! " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

	rabbitmqRes.Destoryy()
	fmt.Println("发送成功！")
}

// RecivMQ 消费队列
func RecivMQ() {
	rabbitmqRes := rabbitmq.NewRabbitMQSimple("golangSimple")
	rabbitmqRes.ConsumeSimple()

	rabbitmqRes.Destoryy()
	fmt.Println("消费成功！")
}

// ReceiveWork work模式消费
func ReceiveWork() {
	rabbitmqRes := rabbitmq.NewRabbitMQSimple("golangSimple")
	rabbitmqRes.ConsumeSimpleWork()

	rabbitmqRes.Destoryy()
	fmt.Println("消费成功！")
}

// SendSub 生产订阅模式消息
func SendSub() {
	rabbitmqRes := rabbitmq.NewRabbitMQPubSub("product_id")

	for i := 0; i < 10; i++ {
		rabbitmqRes.PublishgPub("订阅模式生产第" + strconv.Itoa(i) + "条数据")
		time.Sleep(time.Second)
	}
	rabbitmqRes.Destoryy()
	fmt.Println("sub 发送成功！")
}

// ReceiveSub 订阅消费
func ReceiveSub() {
	rabbitmqRes := rabbitmq.NewRabbitMQPubSub("product_id")
	rabbitmqRes.ReceiveSub()
	rabbitmqRes.Destoryy()
	fmt.Println("消费成功")
}

// SendRout 路由模式生产消息
func SendRout() {
	rabbitmqOne := rabbitmq.NewRabbitMQRouting("exImooc", "imooc_one")
	rabbitmqTwo := rabbitmq.NewRabbitMQRouting("exImooc", "imooc_tow")

	for i := 0; i < 10; i++ {
		rabbitmqOne.PublishgRouting("hello imooc_one " + strconv.Itoa(i))
		rabbitmqTwo.PublishgRouting("hello imooc_two " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

	rabbitmqOne.Destoryy()
	rabbitmqTwo.Destoryy()
}

//ReceiveRout 路由模式消费
func ReceiveRout(routKey string) {
	rabbitmq := rabbitmq.NewRabbitMQRouting("exImooc", routKey)
	rabbitmq.ReceiveRouting()

	rabbitmq.Destoryy()
	log.Printf("消费成功")
}

// SendTopic 主题模式-发送消息
func SendTopic() {
	rabbitmqOne := rabbitmq.NewRabbitMQTopic("exImoocTopic", "imooc.topic.one")
	rabbitmqTwo := rabbitmq.NewRabbitMQTopic("exImoocTopic", "imooc.topic.two")

	for i := 0; i < 10; i++ {
		rabbitmqOne.PublishgTopic("hello imooc.topic one " + strconv.Itoa(i))
		rabbitmqTwo.PublishgTopic("hello imooc.topic two" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

	rabbitmqOne.Destoryy()
	rabbitmqTwo.Destoryy()
	log.Printf("发送成功")
}

// ReceiveTopic 主题模式-消费队列
func ReceiveTopic(topic string) {
	rabbitmq := rabbitmq.NewRabbitMQTopic("exImoocTopic", topic)
	rabbitmq.ReceiveTopic()

	rabbitmq.Destoryy()
	log.Printf("消费成功")
}
