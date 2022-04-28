package helpers

import (
	"bytes"
	"fmt"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

func GenerateStockExcel(products []models.ProductsModel, ctx *fiber.Ctx) (*bytes.Buffer, error) {

	firstRow := []string{"Title", "Slug", "XXS", "XS", "S", "M", "L", "XL", "XXL", "XXL"}

	f := excelize.NewFile()
	// Create a new worksheet.
	activeSheet := f.NewSheet("Sheet1")
	err := f.SetColWidth("Sheet1", "A", "A", 30)
	if err != nil {
		return nil, err
	}
	err = f.SetColWidth("Sheet1", "B", "B", 15)
	if err != nil {
		return nil, err
	}
	// Set value of a cell.
	for i := range firstRow {
		err := f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(65+i), 1), firstRow[i])
		if err != nil {
			return nil, err
		}
	}
	for i := range products {
		err := f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(65), i+2), products[i].Title)
		if err != nil {
			return nil, err
		}
		err = f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(66), i+2), products[i].Slug)
		if err != nil {
			return nil, err
		}
		for j := range products[i].Sizes {
			err := f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(67+j), i+2), products[i].Sizes[j].Stock)
			if err != nil {
				return nil, err
			}
		}
	}
	// Set the active worksheet of the workbook.
	f.SetActiveSheet(activeSheet)

	return f.WriteToBuffer()
}
