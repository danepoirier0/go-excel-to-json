package main

import (
	"encoding/json"
	"log"

	goexceltojson "github.com/danepoirier0/go-excel-to-json"
)

func main() {
	excelFilePath := "./test_file.xlsx"
	excelJsonData, err := goexceltojson.ReadExcel(excelFilePath)
	if err != nil {
		panic(err)
	}
	jsonDataBytes, err := json.Marshal(excelJsonData)
	if err != nil {
		panic(err)
	}
	log.Println("string(jsonDataBytes)", string(jsonDataBytes))
}
