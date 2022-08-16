package cmd

import (
	"github.com/gookit/color"
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

// 根据顺序依次获取参数
var cityCmd = &cobra.Command{
	Use:   "city",
	Short: "",
	Long:  ``,
	Args:  cobra.ExactArgs(2), // 参数必须传两个
	Run: func(cmd *cobra.Command, args []string) {
		color.Red.Println("Do Something.....")

		// 命令行的参数可以直接从args中获取 传参方式：go run main.go name age。
		color.Green.Printf("args[0]: type:%T   value:%v \n", args[0], args[0])
		color.Cyan.Printf("args[1]: type:%T   value:%v", args[1], args[1])
	},
}
