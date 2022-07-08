package cmd

import (
	"CobraApp/app/controllers"
	"CobraApp/app/models/user"
	"CobraApp/pkg/config"
	"CobraApp/pkg/redis"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cityCmd)
}

//noArgs: 如果命令有参数就会报错。
//MaximumNArgs(int)：设置命令可接受的最大参数数量 。
//ExactArgs(int)：如果命令的参数没有指定的数量就会报错。
//RangeArgs(min, max)：命令的参数必须在指定的范围内。

// 运行项目命令 go run main.go city beijing shanghai

var cityCmd = &cobra.Command{
	Use:   "city",
	Short: "",
	Long:  ``,
	Args:  cobra.ExactArgs(2), // 参数必须传两个
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Do Something.....")

		// 命令行的参数可以直接从args中获取 传参方式：go run main.go name age。
		fmt.Println("args[0]: ", args[0])
		fmt.Println("args[1]: ", args[1])

		appName := config.Get("app.name") // 获取配置
		fmt.Println("appName:", appName)

		// logger.ErrorString("Redis", "Decrement", "参数过多") // 记录处理日志方式

		redis.Redis.Set("Cobra_from", "beijing11113", 0) // Redis的使用
		RedisValues := redis.Redis.Get("Cobra_from")
		fmt.Println("RedisValues:", RedisValues)

		userinfo := user.IsEmailExist("11@qq.com") // 查询数据库（mysql）
		fmt.Println("userinfo:", userinfo)

		controllers.GetUserInfo()
	},
}
