package cmd

import (
	"CobraApp/app/controllers"
	"github.com/spf13/cobra"
	"log"
)

// 可以根据参数名称--传参
func init() {
	rootCmd.AddCommand(rabbitmqCmd)
}

// 运行项目命令 go run main.go rabbitmq

var rabbitmqCmd = &cobra.Command{
	Use:     "rabbitmq",
	Short:   "",
	Long:    ``,
	Example: "go run main.go rabbitmq",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			log.Printf("参数错误")
			return
		}

		switch args[0] {
		case "simple": // 简单模式：只有一个消息生产者，一个消息消费者，一个队列，也称为点对点模式、一对一模式
			if args[1] == "send" {
				controllers.SendMQ() //写入队列
			}
			if args[1] == "receive" {
				controllers.RecivMQ() //接受队列信息
			}
			break
		case "work":
			if args[1] == "send" {
				controllers.SendMQ() //写入队列
			}
			if args[1] == "receive" {
				controllers.ReceiveWork() //接受队列信息
			}
			break
		case "sub": //订阅广播模式:一个消息可以同时被多个消费者消费
			if args[1] == "send" {
				controllers.SendSub() //写入队列
			}
			if args[1] == "receive" {
				controllers.ReceiveSub() //接受队列信息
			}
			break
		case "rout": // 路由模式：大体上跟发布订阅模式一样，区别在于发布订阅模式将消息转发给所有绑定的队列，而路由模式将消息转发给那个队列是根据路由匹配情况决定的
			// go run main.go rabbitmq rout receive imooc_tow
			if args[1] == "send" {
				controllers.SendRout()
			}
			if args[1] == "receive" {
				controllers.ReceiveRout(args[2])
			}
			break
		case "topic": // 主题模式：跟路由模式类似，区别在于主题模式的路由匹配支持通配符模糊匹配，而路由模式仅支持完全匹配
			// go run main.go rabbitmq topic receive "imooc.*.two"
			// go run main.go rabbitmq topic receive #
			if args[1] == "send" {
				controllers.SendTopic()
			}
			if args[1] == "receive" {
				controllers.ReceiveTopic(args[2])
			}
			break
		}

	},
}
