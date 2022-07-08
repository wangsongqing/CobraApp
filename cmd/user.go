package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

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
		fmt.Println("name: ", _name)
		fmt.Println("age: ", age)
		fmt.Println("isAge: ", isAge)
	},
}
