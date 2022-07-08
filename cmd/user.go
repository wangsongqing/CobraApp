package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.Flags().StringP("author", "u", "wsq", "姓名")
	userCmd.Flags().StringP("age", "a", "20", "年龄")
}

// 运行项目命令 go run main.go user -u miaowing -a 20
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Do Something.....")

		_name, _ := cmd.Flags().GetString("author")
		age, _ := cmd.Flags().GetString("age")
		fmt.Println("name: ", _name)
		fmt.Println("age: ", age)
	},
}
