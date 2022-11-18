package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

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

// FileExist Exists 判断文件是否存在
func FileExist(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}
