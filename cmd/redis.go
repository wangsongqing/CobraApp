package cmd

import (
	"CobraApp/app/controllers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(RedisCmd)
}

// 运行项目命令 go run main.go test push

// RedisCmd 根据顺序依次获取参数
var RedisCmd = &cobra.Command{
	Use:     "redis",
	Short:   "",
	Long:    ``,
	Example: "go run main.go redis", // 调用实例
	Args:    cobra.ExactArgs(1),     // 参数必须传两个
	Run: func(cmd *cobra.Command, args []string) {

		// initType := []string{"lock"}
		// initName := args[0]
		redisTest := controllers.RedisTest{}
		//redisTest.Lock()
		redisTest.LockApplication()
		//fmt.Println(initName)

	},
}
