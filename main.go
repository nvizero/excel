package main

import (
	"excl/db"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	db.Read()
}

func write() {

	f := excelize.NewFile()
	// 创建一个工作表
	index, _ := f.NewSheet("Sheet2")
	// 设置单元格的值
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

}
