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

// TimeCmd 运行项目命令 go run main.go time
var TimeCmd = &cobra.Command{
	Use:     "time",
	Short:   "",
	Long:    ``,
	Example: "go run main.go time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(time.NowTimeDate())                       // 当前时间:2022-11-17 16:39:07
		fmt.Println(time.GetNowTime())                        // 当前时间:1668674529
		fmt.Println(time.GetTimeNowDate())                    // 只获取年月日
		fmt.Println(time.TimesToStamp("2022-11-12 12:03:23")) // 时间 to 时间戳
		fmt.Println(time.StampToTime(1668736120))             // 时间戳 to 时间
	},
}
