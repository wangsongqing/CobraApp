package cmd

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// 可以根据参数名称--传参
func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.Flags().StringP("author", "u", "miaowing", "姓名")
	userCmd.Flags().Int64P("age", "a", 20, "年龄")
	userCmd.Flags().BoolP("isShow", "i", false, "是否显示年龄")
}

// 运行项目命令 go run main.go user -u miaowing -a 20 -i true
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Do Something.....")

		_name, _ := cmd.Flags().GetString("author") // 获取字符串参数
		age, _ := cmd.Flags().GetInt64("age")       // 获取int类型参数
		isAge, _ := cmd.Flags().GetBool("isShow")   // 获取bool类型参数
		color.Red.Println("name: ", _name)
		color.Green.Println("age: ", age)
		color.Cyan.Println("isAge: ", isAge)
	},
}
