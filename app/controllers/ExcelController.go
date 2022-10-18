package controllers

import (
	"CobraApp/app/models/links"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type ExcelTest struct {
	ExcelPath string
}

// 读取Excel数据
func (e *ExcelTest) Read() {
	f, err := excelize.OpenFile(e.ExcelPath + "Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		fmt.Println(row[1])
	}
}

// 数据写入Excel
func (e *ExcelTest) Write() {
	result := links.GetLink()
	f := excelize.NewFile()

	// 创建一个工作表
	index := f.NewSheet("Sheet1")

	// 设置单元格的值
	for key, value := range result {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(key), value.Name)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(key), value.Url)
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)

	// 根据指定路径保存文件
	if err := f.SaveAs(e.ExcelPath + "Book1.xlsx"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Success")
}
