package utils

import (
	"log"

	"github.com/hosseinmirzapur/parsian-backend/data/models"
	"github.com/xuri/excelize/v2"
)

func ExcelExport(data []models.OrderItem) (*excelize.File, error) {
	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("error closing excel file: %v", err)
		}
	}()

	// Sheet initialization
	sheet := "Sheet1"
	index, err := f.NewSheet(sheet)
	if err != nil {
		return nil, err
	}
	f.SetActiveSheet(index)

	// Set Headers
	f.SetCellValue(sheet, "A1", "نام")
	f.SetCellValue(sheet, "B1", "اجازه سنباده")
	f.SetCellValue(sheet, "C1", "اجازه تخریب")
	f.SetCellValue(sheet, "D1", "وضعیت")
	f.SetCellValue(sheet, "E1", "نوع آزمون")
	f.SetCellValue(sheet, "F1", "تعداد")
	f.SetCellValue(sheet, "G1", "لینک عکس")

	// Set Data
	for i, item := range data {
		f.SetCellValue(sheet, getCell(i+2, 1), item.Name)
		f.SetCellValue(sheet, getCell(i+2, 2), handlePermissions(item.AllowSandPaper))
		f.SetCellValue(sheet, getCell(i+2, 3), handlePermissions(item.AllowDestruction))
		f.SetCellValue(sheet, getCell(i+2, 4), handleStatus(item.Status))
		f.SetCellValue(sheet, getCell(i+2, 5), handleTestType(item.TestType))
		f.SetCellValue(sheet, getCell(i+2, 6), item.Quantity)
		f.SetCellValue(sheet, getCell(i+2, 7), item.FilePath)
	}

	// Return File
	return f, nil

}

func handlePermissions(allowed bool) string {
	if allowed {
		return "دارد"
	} else {
		return "ندارد"
	}
}

func handleStatus(status string) string {
	switch status {
	case "pending":
		return "در حال بررسی"
	case "partial":
		return "پرداخت جزئی"
	case "office":
		return "حساب دفتری"
	case "paid":
		return "پرداخت شده"
	default:
		return ""
	}
}

func handleTestType(testType string) string {
	switch testType {
	case "analyze":
		return "آنالیز"
	case "hardness":
		return "سختی"
	case "both":
		return "هر دو"
	default:
		return ""
	}
}

func getCell(row, col int) string {
	address, err := excelize.CoordinatesToCellName(col, row)
	if err != nil {
		return ""
	}

	return address
}
