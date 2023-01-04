package cmd

import (
	"CobraApp/app/controllers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(MysqlCmd)
}

// 运行项目命令 go run main.go redis

// MysqlCmd redis分布式锁
var MysqlCmd = &cobra.Command{
	Use:     "mysql",
	Short:   "",
	Long:    ``,
	Example: "go run main.go mysql", // 调用实例
	//Args:    cobra.ExactArgs(2),     // 参数必须传两个
	Run: func(cmd *cobra.Command, args []string) {

		mysqlController := controllers.MysqlController{}
		mysqlController.GetUserInfo()
		//mysqlController.GetUserListById()
		//mysqlController.Create()
		//mysqlController.Update()
		//mysqlController.Updates()
		//mysqlController.Transaction()
		//mysqlController.Transaction01()
		//mysqlController.Delete()
		//mysqlController.Deletes()
		//mysqlController.Creates()
		//mysqlController.Page()
	},
}
