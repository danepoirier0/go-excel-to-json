package tests

import (
	"encoding/json"
	"testing"

	goexceltojson "github.com/danepoirier0/go-excel-to-json"
)

func TestReadExcel(t *testing.T) {
	excelFilePath := "../test_file.xlsx"
	excelJsonData, err := goexceltojson.ReadExcel(excelFilePath)
	if err != nil {
		panic(err)
	}
	jsonDataBytes, err := json.Marshal(excelJsonData)
	if err != nil {
		panic(err)
	}
	t.Log("string(jsonDataBytes)", string(jsonDataBytes))
}

func TestReadExcelToJsonString(t *testing.T) {
	excelFilePath := "../test_file.xlsx"
	sheetIndex := 0
	pageSize := 3
	pageIndex := 2
	sheetRows, jsonString, err := goexceltojson.ReadExcelToPagedJsonString(excelFilePath, sheetIndex, pageSize, pageIndex)
	if err != nil {
		panic(err)
	}
	t.Log("sheetRows", sheetRows)
	t.Log("jsonString", jsonString)
}
