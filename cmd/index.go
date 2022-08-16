package cmd

import (
	"CobraApp/app/controllers"
	"CobraApp/app/models/user"
	"CobraApp/pkg/config"
	"CobraApp/pkg/redis"
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// 可以根据参数名称--传参
func init() {
	rootCmd.AddCommand(indexCmd)
	indexCmd.Flags().Int64P("id", "i", 0, "USER_ID")
	indexCmd.Flags().StringP("email", "e", "", "EMAIL")
}

//noArgs: 如果命令有参数就会报错。
//MaximumNArgs(int)：设置命令可接受的最大参数数量 。
//ExactArgs(int)：如果命令的参数没有指定的数量就会报错。
//RangeArgs(min, max)：命令的参数必须在指定的范围内。

// 运行项目命令 go run main.go index -i 1 -e 11@qq.com

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		appName := config.Get("app.name") // 获取配置
		color.Green.Println("appName:", appName)
		fmt.Println("---------------------------------------------------------------------------")

		// logger.ErrorString("Redis", "Decrement", "参数过多") // 记录处理日志方式

		redis.Redis.Set("Cobra_from", "this is a good man", 0) // Redis的使用
		RedisValues := redis.Redis.Get("Cobra_from")
		color.Cyan.Println("RedisValues:", RedisValues)
		fmt.Println("---------------------------------------------------------------------------")

		email, _ := cmd.Flags().GetString("email")
		userinfo := user.IsEmailExist(email) // 查询数据库（mysql）
		color.Yellow.Println("RedisValues:", userinfo)
		fmt.Println("---------------------------------------------------------------------------")

		id, _ := cmd.Flags().GetInt64("id")
		controllers.GetUserInfo(id)
	},
}
