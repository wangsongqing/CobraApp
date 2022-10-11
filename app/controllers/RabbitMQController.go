package controllers

import (
	"CobraApp/pkg/rabbitmq"
	"fmt"
	"strconv"
	"time"
)

// SendMQ 写入队列
func SendMQ() {
	rabbitmqRes := rabbitmq.NewRabbitMQSimple("imoocSimple")
	rabbitmqRes.PublishSimple("Hello imooc!")
	rabbitmqRes.Destoryy()
	fmt.Println("发送成功！")
}

// RecivMQ 消费队列
func RecivMQ() {
	rabbitmqRes := rabbitmq.NewRabbitMQSimple("imoocSimple")
	rabbitmqRes.ConsumeSimple()
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
}

//ReceiveRout 路由模式消费
func ReceiveRout(routKey string) {
	rabbitmqOne := rabbitmq.NewRabbitMQRouting("exImooc", routKey)
	rabbitmqOne.ReceiveRouting()
}
