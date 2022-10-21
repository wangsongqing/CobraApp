package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
	"github.com/parnurzeal/gorequest"
)

type Request struct {
	Url string
}

func (r *Request) Post(param map[string]string) (int, map[string]interface{}, []error) {
	// post json方式请求
	request := gorequest.New()

	resp, body, errs := request.Post(r.Url).
		Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjA3NjIxMmNkMzY2NzFmOWZkNGQyMjkwOGI0NTA1NjJkZGFjMDVjZmI4ZTFkOWJlOTk5OWY3Y2Y4MDgxMGJlYWM5NjBiNGU4ZDUyZDlmNmQ3In0.eyJhdWQiOiIyIiwianRpIjoiMDc2MjEyY2QzNjY3MWY5ZmQ0ZDIyOTA4YjQ1MDU2MmRkYWMwNWNmYjhlMWQ5YmU5OTk5ZjdjZjgwODEwYmVhYzk2MGI0ZThkNTJkOWY2ZDciLCJpYXQiOjE2NjYxNjEzMzYsIm5iZiI6MTY2NjE2MTMzNiwiZXhwIjoxNjY4NzUzMzM2LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JZUJOxtIq38X4kTECMfyQGywBW43JQhqE9aPucOqLfire-civf3AiQzPr3zhpwwpMTqpW95rk3Rs0zk3ejKDQX7pah7sh0ljFLVg7AN4XhRVex-_0sfEev50m_0GmCjEkG4WC2TvMFDfjGGBPbUU8ytEm_DE_YNcDhM29G4CRgihiTKxqsFIW9HckieEKFfcd9yvl_DMajEM1qqeJQWT02eMrnN9MpoBcVGfi1lNwpMTN7hxaKTxWRA6bZQHjYP4YXQKvGxyXWPdqE0ZvLxV6XWoiNRQnLY8BlSMsb0DHNsglaIRjVpKBGymVqVGx3NJvcXNzI7v3Siud4Fb-AtBUJZDqBdMZygwtMs54acNE23cVnCZ2hEWB_pnGkK_YPfm4d9GTLUipYHtRdKA1vPIulu42ROlFpV9FCLNADgyJR0dObN2RpC0JoGk2g78U7oRGNsF2B7p0R7gKyHKgyYilE924Zo9KuZ0G9zG0T9EyxDX1RJyNC_meUvg-ADoWvn8aX3usDZawsPvDONwQXNpv8UL_Z863zfTffdExHaB9rdUOUeA0HXHPEB7vAWNVGJvTZC8rzYXkBy1_P21duF-vmL0sywxslY35uflQCj_cWpXOLwXJ0XT94_o4Eqm2FcemuKt7wsoBVQDqL-tvcw-Uu4XmoJLC3d7aMnTaRTKNrY").
		Send(param).
		EndBytes()

	data := make(map[string]interface{})
	err := json.Unmarshal(body, &data)

	if err != nil {
		return 0, nil, nil
	}

	return resp.StatusCode, data, errs
}

func (r *Request) GetMap(param map[string]string, data map[string]interface{}) (int, []error) {
	// get json方式请求
	request := gorequest.New()

	resp, body, errs := request.Get(r.Url).
		Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjA3NjIxMmNkMzY2NzFmOWZkNGQyMjkwOGI0NTA1NjJkZGFjMDVjZmI4ZTFkOWJlOTk5OWY3Y2Y4MDgxMGJlYWM5NjBiNGU4ZDUyZDlmNmQ3In0.eyJhdWQiOiIyIiwianRpIjoiMDc2MjEyY2QzNjY3MWY5ZmQ0ZDIyOTA4YjQ1MDU2MmRkYWMwNWNmYjhlMWQ5YmU5OTk5ZjdjZjgwODEwYmVhYzk2MGI0ZThkNTJkOWY2ZDciLCJpYXQiOjE2NjYxNjEzMzYsIm5iZiI6MTY2NjE2MTMzNiwiZXhwIjoxNjY4NzUzMzM2LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JZUJOxtIq38X4kTECMfyQGywBW43JQhqE9aPucOqLfire-civf3AiQzPr3zhpwwpMTqpW95rk3Rs0zk3ejKDQX7pah7sh0ljFLVg7AN4XhRVex-_0sfEev50m_0GmCjEkG4WC2TvMFDfjGGBPbUU8ytEm_DE_YNcDhM29G4CRgihiTKxqsFIW9HckieEKFfcd9yvl_DMajEM1qqeJQWT02eMrnN9MpoBcVGfi1lNwpMTN7hxaKTxWRA6bZQHjYP4YXQKvGxyXWPdqE0ZvLxV6XWoiNRQnLY8BlSMsb0DHNsglaIRjVpKBGymVqVGx3NJvcXNzI7v3Siud4Fb-AtBUJZDqBdMZygwtMs54acNE23cVnCZ2hEWB_pnGkK_YPfm4d9GTLUipYHtRdKA1vPIulu42ROlFpV9FCLNADgyJR0dObN2RpC0JoGk2g78U7oRGNsF2B7p0R7gKyHKgyYilE924Zo9KuZ0G9zG0T9EyxDX1RJyNC_meUvg-ADoWvn8aX3usDZawsPvDONwQXNpv8UL_Z863zfTffdExHaB9rdUOUeA0HXHPEB7vAWNVGJvTZC8rzYXkBy1_P21duF-vmL0sywxslY35uflQCj_cWpXOLwXJ0XT94_o4Eqm2FcemuKt7wsoBVQDqL-tvcw-Uu4XmoJLC3d7aMnTaRTKNrY").
		Send(param).
		EndBytes()

	err := json.Unmarshal(body, &data)

	if err != nil {
		return 0, nil
	}

	return resp.StatusCode, errs
}

func (r *Request) GetStruct(param interface{}, data any) {
	resp, err := requests.Get(r.Url, param)
	if err != nil {
		return
	}
	// 状态码
	fmt.Println(resp.R.StatusCode)
	// 响应体
	// println(resp.Text())

	//resStruct := struct {
	//	Aaa int `json:"aaa"`
	//}{}
	json.Unmarshal(resp.Content(), &data)
	fmt.Println(data)
}
