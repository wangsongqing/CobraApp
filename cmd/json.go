package cmd

import (
	"CobraApp/app/controllers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(JsonCmd)
}

//noArgs: 如果命令有参数就会报错。
//MaximumNArgs(int)：设置命令可接受的最大参数数量 。
//ExactArgs(int)：如果命令的参数没有指定的数量就会报错。
//RangeArgs(min, max)：命令的参数必须在指定的范围内。

// 运行项目命令 go run main.go json push

// JsonCmd TestCmd 根据顺序依次获取参数
var JsonCmd = &cobra.Command{
	Use:     "json",
	Short:   "",
	Long:    ``,
	Example: "go run main.go json (decode || encode)", // 调用实例
	Run: func(cmd *cobra.Command, args []string) {
		jsonTest := controllers.JsonTest{}

		if args[0] == "encode" {
			jsonTest.JsonEncode()
		}

		if args[0] == "encode_list" {
			jsonTest.JsonEncodeList()
		}

		if args[0] == "decode" {
			jsonTest.JsonDecode()
		}

		if args[0] == "decode_list" {
			jsonTest.JsonDecodeList()
		}

		if args[0] == "gjson" { // gjson包的使用
			jsonTest.GjsonTest()
		}

	},
}
