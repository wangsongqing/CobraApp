package cmd

import (
	"CobraApp/pkg/time"
	"fmt"
	"github.com/spf13/cobra"
)

// 可以根据参数名称--传参
func init() {
	rootCmd.AddCommand(TimeCmd)
}

var TimeCmd = &cobra.Command{
	Use:     "time",
	Short:   "",
	Long:    ``,
	Example: "go run main.go time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(time.NowTimeDate())    // 当前时间:2022-11-17 16:39:07
		fmt.Println(time.GetNowTime())     // 当前时间:1668674529
		fmt.Println(time.GetTimeNowDate()) // 只获取年月日
	},
}
