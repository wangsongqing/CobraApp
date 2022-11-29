package cmd

import (
	"CobraApp/app/controllers"
	"github.com/spf13/cobra"
)

// 可以根据参数名称--传参
func init() {
	rootCmd.AddCommand(ElasticsearchCmd)
}

//运行项目命令 go run main.go es

var ElasticsearchCmd = &cobra.Command{
	Use:     "es",
	Short:   "",
	Long:    ``,
	Example: "go run main.go es (create || select || update || delete)", // 调用实例
	Run: func(cmd *cobra.Command, args []string) {
		argsType := args[0]
		es := controllers.Elasticsearch{}
		if argsType == "create" { // 创建索引
			es.Create()
		}

		if argsType == "add" {
			es.Add()
		}

		if argsType == "search" {
			es.Search()
		}

		if argsType == "update" {
			es.Update()
		}

	},
}
