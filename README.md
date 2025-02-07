### go读取excel转换成json数据


##### 获取依赖```go get github.com/danepoirier0/go-excel-to-json@v0.0.1```

##### 示例
```go
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


```