package cmd

import (
	"CobraApp/app/controllers"
	"CobraApp/pkg/helpers"
	"github.com/gookit/color"
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
	Example: "go run main.go es (create || select || update || delete || add || search)", // 调用实例
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			color.Redln("参数错误")
			return
		}

		argsType := args[0]

		types := []string{"create", "select", "update", "delete", "add", "search"}
		flag := helpers.SliceInString(types, argsType)
		if flag == false {
			color.Redln("参数错误，请输入 (create || select || update || delete || add) 类型")
			return
		}

		es := controllers.Elasticsearch{}
		if argsType == "create" { // 创建索引
			es.Create()
		}

		// 添加文档
		if argsType == "add" {
			es.Add()
		}

		// 查询文档
		if argsType == "search" {
			es.Search()
		}

		// 更新文档
		if argsType == "update" {
			es.Update()
		}

		// 删除文档
		if argsType == "delete" {
			es.Delete()
		}
	},
}
