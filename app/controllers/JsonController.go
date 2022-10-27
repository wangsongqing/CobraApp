package controllers

import (
	"CobraApp/pkg/helpers"
	"encoding/json"
	"fmt"
)

type JsonTest struct {
	Types string `json:"types"`
}

type UserTest struct {
	Address struct {
		City string `json:"city"`
		Home string `json:"home"`
	} `json:"Address"`
	Age  string `json:"age"`
	Name string `json:"name"`
}

// JsonEncode JsonEncpde 相当于PHP的json_encode
func (js *JsonTest) JsonEncode() {
	map1 := map[string]string{
		"name": "123",
	}

	data, _ := helpers.JsonEncodeToMap(map1)
	fmt.Printf("type:%T, data: %v", string(data), string(data))
	// 执行结果 type:string, data: {"name":"123"}
}

func (js *JsonTest) JsonEncodeList() {
	test := &JsonTest{Types: "abc"}
	data, _ := json.Marshal(test)
	fmt.Printf("type:%T, data: %v", string(data), string(data))

	// 执行结果 type:string, data: {"types":"abc"}
}

// JsonDecode 一为数组json
func (js *JsonTest) JsonDecode() {
	//data := links.Get(7)
	jsonStr := `{"name":"xxx"}`
	var url map[string]interface{}
	urlData, _ := helpers.JsonDecodeToMap(helpers.String2Bytes(jsonStr), url)

	for k, v := range urlData {
		fmt.Printf("type:%v, data:%v \n", k, v)
	}

	// 执行结果 type:name, data:xxx
}

// JsonDecodeList 多维数字json
func (js *JsonTest) JsonDecodeList() {
	var user []UserTest
	byteData := `[{"name":"xxx","age":"18","Address":{"city":"beijing","home":"haidian"}},{"name":"aaaa","age":"20","Address":{"city":"chengdu","home":"shaungliu"}}]`
	if err := json.Unmarshal([]byte(byteData), &user); err != nil {
		return
	}

	for k, v := range user {
		fmt.Printf("k : %+v, v : %+v \n", k, v.Address.City)
	}

	// 输入结果
	// k : 0, v : beijing
	// k : 1, v : chengdu
}
