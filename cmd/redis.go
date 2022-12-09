package cmd

import (
	"CobraApp/app/controllers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(RedisCmd)
}

// 运行项目命令 go run main.go redis

// RedisCmd redis分布式锁
var RedisCmd = &cobra.Command{
	Use:     "redis",
	Short:   "",
	Long:    ``,
	Example: "go run main.go redis", // 调用实例
	//Args:    cobra.ExactArgs(2),     // 参数必须传两个
	Run: func(cmd *cobra.Command, args []string) {

		redisTest := controllers.RedisTest{}
		redisTest.Lock()

	},
}
