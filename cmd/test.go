package cmd

import (
	"CobraApp/app/controllers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(TestCmd)
}

//noArgs: 如果命令有参数就会报错。
//MaximumNArgs(int)：设置命令可接受的最大参数数量 。
//ExactArgs(int)：如果命令的参数没有指定的数量就会报错。
//RangeArgs(min, max)：命令的参数必须在指定的范围内。

// 运行项目命令 go run main.go city beijing shanghai

// TestCmd 根据顺序依次获取参数
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "",
	Long:  ``,
	//Args:  cobra.ExactArgs(2), // 参数必须传两个
	Run: func(cmd *cobra.Command, args []string) {
		// color.Red.Println("Do Something.....")

		// 命令行的参数可以直接从args中获取 传参方式：go run main.go name age。
		// color.Green.Printf("args[0]: type:%T   value:%v \n", args[0], args[0])
		// color.Cyan.Printf("args[1]: type:%T   value:%v", args[1], args[1])

		initType := args[0]
		var test = controllers.Test{}
		if initType == "push" {
			test.ChannelPushList()
		}

		if initType == "pop" {
			test.ChannelPopList()
		}
	},
}
