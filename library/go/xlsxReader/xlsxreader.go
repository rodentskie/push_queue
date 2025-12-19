package xlsxreader

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// SheetData holds the sheet name and its rows
// Each row is a slice of strings
// You may want to extend this for more complex logic
// For now, it just returns all rows as [][]string

type SheetData struct {
	Name string
	Rows [][]string
}

// ReadXLSX reads all sheets and returns their data
func ReadXLSX(path string) ([]SheetData, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open xlsx: %w", err)
	}
	var sheets []SheetData
	for _, name := range f.GetSheetList() {
		rows, err := f.GetRows(name)
		if err != nil {
			return nil, fmt.Errorf("failed to read sheet %s: %w", name, err)
		}
		sheets = append(sheets, SheetData{Name: name, Rows: rows})
	}
	return sheets, nil
}
