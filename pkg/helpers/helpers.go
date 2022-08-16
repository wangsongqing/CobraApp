// Package helpers 存放辅助方法
package helpers

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

// ReadFile 按行读取文件
func ReadFile(fileName string) ([]string, error) {
	var fileData []string
	files, err := os.Open(fileName)
	if err != nil {
		return fileData, err
	}

	defer func(files *os.File) {
		err := files.Close()
		if err != nil {
			return
		}
	}(files)

	render := bufio.NewReader(files)

	for {
		str, err := render.ReadString('\n')
		fileData = append(fileData, fmt.Sprintf("%s", str))
		if err == io.EOF { // io.EOF 表示文件末行
			break
		}
	}

	return fileData, nil
}

// ReadAll 一次性读取文件到内存
func ReadAll(fileName string) string {
	contentList, err := ioutil.ReadFile(fileName)

	if err != nil {
		return ""
	}

	return string(contentList)
}

// WriteFile 追加方式写入文件,如果文件不存在则新建文件
func WriteFile(fileName string, content string) (bool, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return false, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	write := bufio.NewWriter(file)
	writeString, err := write.WriteString(content + "\r\n")
	if err != nil {
		return false, err
	}

	if writeString < 0 {
		return false, err
	}

	err = write.Flush()
	if err != nil {
		return false, err
	}

	return true, nil
}

// FileExist 判断文件是否存在
func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)

	var isExit bool
	if err == nil {
		isExit = true
	}

	if os.IsNotExist(err) {
		isExit = false
	}

	return isExit
}
