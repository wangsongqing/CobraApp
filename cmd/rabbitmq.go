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
	Use:   "rabbitmq",
	Short: "",
	Long:  ``,
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
		case "sub": //订阅广播模式:一个消息可以同时被多个消费者消费
			if args[1] == "send" {
				controllers.SendSub() //写入队列
			}
			if args[1] == "receive" {
				controllers.ReceiveSub() //接受队列信息
			}
			break
		}

	},
}
