package cmd

import (
	"CobraApp/app/controllers"
	"CobraApp/pkg/logger"
	requests "CobraApp/pkg/request"
	"fmt"
	"github.com/spf13/cobra"
)

type EventName struct {
	ErrorCode int64  `json:"error_code"`
	Reason    string `json:"reason"`
	Result    []struct {
		Date  string `json:"date"`
		Day   string `json:"day"`
		EID   string `json:"e_id"`
		Title string `json:"title"`
	} `json:"result"`
}

// 可以根据参数名称--传参
func init() {
	rootCmd.AddCommand(RequestCmd)
}

// RequestCmd 运行项目命令 go run main.go request post
// 安装库命令: go get github.com/parnurzeal/gorequest
var RequestCmd = &cobra.Command{
	Use:     "request",
	Short:   "",
	Long:    ``,
	Example: "go run main.go request (post || get)",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) <= 0 {
			fmt.Println("参数错误")
			return
		}
		handleType := args[0]
		request := controllers.Request{}

		switch handleType {
		case "post":
			request.Url = "http://shopifyerp.org/api/logs/import_product?page_size=2"
			param := map[string]string{"name": "backy", "species": "dog"} // 传参
			resp, body, errs := request.Post(param)

			if resp != 200 || len(errs) > 0 {
				fmt.Printf("请求错误, resp:%v,body:%v, errs:%v", resp, body, errs)
				return
			}

			fmt.Printf("body:%v \n", body)
			fmt.Printf("errs:%v \n", errs)
			break
		case "get":
			request.Url = "http://shopifyerp.org/api/custom_skus?filter%5BcustomProduct.id%5D=3483"
			param := map[string]string{"page": "1", "page_size": "2"} // 传参
			result := make(map[string]interface{})
			resp, errs := request.GetMap(param, result)

			if resp != 200 || len(errs) > 0 {
				fmt.Printf("请求错误, resp:%v,body:%v, errs:%v", resp, result, errs)
				return
			}

			data := result["data"].([]interface{})
			for k, v := range data {
				value := v.(map[string]interface{}) // interface转map
				fmt.Printf("key:%v, value:%v \n", k, value["id"])
			}

			//fmt.Printf("type:%T,body:%v \n", body["data"], body["data"])
			//fmt.Printf("errs:%v \n", errs)
			break

		case "event":
			url := "http://v.juhe.cn/todayOnhistory/queryEvent.php?date=1/1&key=xxx"
			data := map[string]string{"date": "1/1", "key": "245fb61abff784de7e15a6752e2c84b6"}
			headers := map[string]string{"Authorization": "6c38874672f28b8a11"}
			request := requests.Request{Url: &url, Method: "GET", Data: &data, Headers: &headers}
			request.Body()
			list := request.RepInfo

			if _, ok := list["result"]; ok || list["error_code"] != 0 {
				log := fmt.Sprintf("请求异常，异常原因:%v, url:%v, parama:%v", list["reason"], url, data)
				logger.Info(log)
				break
			}

			// 遍历map
			result := list["result"].([]interface{})
			for _, value := range result {
				values := value.(map[string]interface{})
				fmt.Printf("day:%v, date:%v, title:%v, e_id:%v \n", values["day"], values["date"], values["title"], values["e_id"])
			}
			break

		}
	},
}
