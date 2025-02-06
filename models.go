package goexceltojson

type ExcelJsonSheet struct {
	SheetName string   // sheet名称
	Rows      int      // 当前sheet行数
	Headers   []string // 表头
	Data      []map[string]string
}

type ExcelJsonData struct {
	FileName  string // 文件名
	TotalRows int    // 所有sheets的总行数
	Sheets    []ExcelJsonSheet
}
