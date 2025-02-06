package goexceltojson

import (
	"errors"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

func ReadExcel(excelFilepath string) (*ExcelJsonData, error) {
	fileName, fileExt, err := getFileInfoFromPath(excelFilepath)
	if err != nil {
		return nil, err
	}

	if fileExt != ".xlsx" {
		return nil, errors.New("only support xlsx file")
	}

	rtData := &ExcelJsonData{
		FileName: fileName,
	}

	// 打开 Excel 文件
	f, err := excelize.OpenFile(excelFilepath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	sheetNames := f.GetSheetList()
	// 遍历每个工作表, 并获取数据
	for _, sheetName := range sheetNames {
		// 获取工作表的行数
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return nil, err
		}
		// 获取表头
		headers := rows[0]
		data := make([]map[string]string, 0)
		// 遍历数据行
		for _, row := range rows[1:] {
			if len(row) == 0 {
				continue
			}
			// log.Println("row", row)
			// log.Println("len(row)", len(row))
			rowData := make(map[string]string)

			// 遍历headers对应的列
			for i, cell := range headers {
				if i >= len(row) {
					continue
				}
				rowData[cell] = row[i]
			}
			data = append(data, rowData)
		}

		rtData.Sheets = append(rtData.Sheets, ExcelJsonSheet{
			SheetName: sheetName,
			Rows:      len(rows),
			Headers:   headers,
			Data:      data,
		})
	}

	return rtData, nil
}

func ReadExcelToPagedJsonString(excelFilepath string, sheetIndex, pageSize, pageIndex int) (string, error) {
	excelData, err := ReadExcel(excelFilepath)
	if err != nil {
		return "", err
	}
	if sheetIndex >= len(excelData.Sheets) {
		return "", errors.New("sheetIndex out of range")
	}

	// TODO
	return "", errors.New("not implement")
}

// 获取文件名称和扩展名
func getFileInfoFromPath(filePath string) (string, string, error) {
	fileName := filepath.Base(filePath) // 带扩展名
	ext := filepath.Ext(filePath)

	return fileName, ext, nil
}
