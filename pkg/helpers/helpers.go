// Package helpers 存放辅助方法
package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"time"
)

// Empty 类似于 PHP 的 empty() 函数
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// NowTime 当前时间
func NowTime() string {
	now := time.Now()
	// return fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	return fmt.Sprintf("date:%v", now.Format("2006-01-02 15:04:05"))
}

// String2Bytes string 转换成byte
func String2Bytes(data string) []byte {
	return []byte(data)
}

// FmtStrFromInterface 转换interface到字符串
func FmtStrFromInterface(val interface{}) string {
	if val == nil {
		return ""
	}
	switch ret := val.(type) {
	case string:
		return ret
	case int8, uint8, int16, uint16, int, uint, int64, uint64, float32, float64:
		return fmt.Sprintf("%v", ret)
	}
	return ""
}

// JsonToMap JSON转map类型 /**
func JsonToMap(body []byte) map[string]interface{} {
	mapList := make(map[string]interface{})
	errors := json.Unmarshal(body, &mapList)
	if errors != nil {
		fmt.Println(errors)
		return mapList
	}

	return mapList
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// MakeRandInt 随机数
func MakeRandInt() int {
	min := 1000
	max := 9999
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(max-min) + min

	return number
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// JsonEncodeToMap JsonEncode Json_encode 类型转为json ，配合JsonToMap使用效果更好
func JsonEncodeToMap(data any) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	return jsonData, err
}

func JsonDecodeToMap(byteData []byte, data map[string]interface{}) (map[string]interface{}, error) {
	err := json.Unmarshal(byteData, &data)
	return data, err
}

// SliceInString SliceInData 检查数据是否在slice
func SliceInString(slices []string, value string) bool {
	flag := false

	for _, v := range slices {
		if value == v {
			flag = true
		}
	}

	return flag
}
