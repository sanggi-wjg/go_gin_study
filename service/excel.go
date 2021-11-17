package service

import (
	"go_gin_study/server"

	"github.com/xuri/excelize"
)

func CreateExcel() {
	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large",
		"B1": "Apple", "C1": "Orange", "D1": "Pear"}

	file := excelize.NewFile()
	for k, v := range categories {
		file.SetCellValue("Sheet1", k, v)
	}

	if err := file.SaveAs("test.xlsx"); err != nil {
		server.Error(err)
	}
}
