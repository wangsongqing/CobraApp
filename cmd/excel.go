package cmd

import (
	"CobraApp/app/controllers"
	"fmt"
	"github.com/spf13/cobra"
)

// 可以根据参数名称--传参
func init() {
	rootCmd.AddCommand(ExcelCmd)
}

//运行项目命令 go run main.go excel

var ExcelCmd = &cobra.Command{
	Use:     "excel",
	Short:   "",
	Long:    ``,
	Example: "go run main.go excel", // 调用实例
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) <= 0 {
			fmt.Println("参数错误")
			return
		}
		handleType := args[0]
		excel := controllers.ExcelTest{}
		excel.ExcelPath = "/Users/songsong/development/Glang/src/github.com/wangsongqing/CobraApp/"
		if handleType == "read" {
			excel.Read()
			return
		}

		// 数据写入Excel
		if handleType == "write" {
			excel.Write()
			return
		}
	},
}
