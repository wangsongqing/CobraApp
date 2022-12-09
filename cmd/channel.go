package cmd

import (
	"CobraApp/app/controllers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ChannelCmd)
}

// 运行项目命令 go run main.go channel push

// ChannelCmd channel的使用
var ChannelCmd = &cobra.Command{
	Use:     "channel",
	Short:   "",
	Long:    ``,
	Example: "go run main.go channel (push | pop | mpop)", // 调用实例
	Args:    cobra.ExactArgs(1),                           // 参数必须传两个
	Run: func(cmd *cobra.Command, args []string) {

		redisKey := "redis:channel:cobra-app"
		redisTest := controllers.ChannelTest{RedisKey: redisKey}
		argsType := args[0]
		// 数据写入队列
		if argsType == "push" {
			redisTest.Push()
		}

		// 单协程消费
		if argsType == "pop" {
			redisTest.Pop()
		}

		// 多协程消费
		if argsType == "mpop" {
			redisTest.MultiProcessingPop()
		}

	},
}
