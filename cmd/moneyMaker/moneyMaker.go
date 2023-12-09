package main

import (
	"fmt"
	"log"
	kpParser "tommy2thicc/internal/parser"
	"tommy2thicc/internal/spreadsheet"

	"github.com/tealeg/xlsx"
)

func main() {

	teamStats, err := kpParser.ParseKenPomHtml()
	if err != nil {
		log.Fatal(err)
	}

	// Open the Excel file
	xlFile, err := xlsx.OpenFile("games.xlsx")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	var rows [][]string

	// Iterate through each sheet in the Excel file
	for _, sheet := range xlFile.Sheets {

		// Iterate through each row in the sheet
		for _, row := range sheet.Rows {
			var rowCells []string

			// Iterate through each cell in the row
			for _, cell := range row.Cells {
				text := cell.String()
				rowCells = append(rowCells, text)
			}

			rows = append(rows, rowCells)
		}
	}
	goldenRows := spreadsheet.AnalyzeTheSheet(rows[1:], teamStats)

	excelFile := xlsx.NewFile()
	sheet, err := excelFile.AddSheet("GoldenBoy")
	if err != nil {
		fmt.Printf("Error creating sheet: %s\n", err)
		return
	}

	headers := []string{"Team", "Expected ML Log5", "Expected ML Kp",
		"Predicted Points Log5", "Spread Kp", "Total Points Log5", "Log5 WP", "Kp WP", "", "Team", "Vegas Spread", "Vegas O/U", "Vegas WP"}

	headerRow := sheet.AddRow()
	for i, header := range headers {
		style := xlsx.NewStyle()
		style.Border.Bottom = "thin"
		style.Font = *xlsx.NewFont(12, "Times New Roman")
		if i == 0 || i == 2 || i == 5 || i == 9 {
			style.Border.Right = "thick"
		}
		cell := headerRow.AddCell()
		cell.Value = header
		cell.SetStyle(style)
		// Set the column width to accommodate the content
		if i != 8 {
			sheet.SetColWidth(i, i, float64(13)) // Set the width for the first column (0-based index)
		} else {
			sheet.SetColWidth(i, i, float64(4)) // Set the width for the first column (0-based index)
		}
	}

	for i, rowData := range goldenRows {
		style := xlsx.NewStyle()
		style.Font = *xlsx.NewFont(11, "Times New Roman")

		styleColumnLine := xlsx.NewStyle()
		styleColumnLine.Font = *xlsx.NewFont(11, "Times New Roman")
		styleColumnLine.Border.Right = "thick"
		if i%2 == 1 {
			style.Border.Bottom = "thin"
			styleColumnLine.Border.Bottom = "thin"
		}
		row := sheet.AddRow()

		cell := row.AddCell()
		cell.Value = rowData.Name
		cell.SetStyle(styleColumnLine)

		cell = row.AddCell()
		cell.SetValue(rowData.ExpectedMoneyLineLog5)
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.SetValue(rowData.ExpectedMoneyLineKp)
		cell.SetStyle(styleColumnLine)

		cell = row.AddCell()
		cell.SetValue(rowData.PredictedPointsLog5)
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.SetValue(rowData.KpSpread)
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.SetValue(rowData.Log5PredictedTotal)
		cell.SetStyle(styleColumnLine)

		cell = row.AddCell()
		cell.SetValue(rowData.WinPercentageLog5)
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.SetValue(rowData.KpWinPercentage)
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.Value = ""

		cell = row.AddCell()
		cell.Value = rowData.Name
		cell.SetStyle(styleColumnLine)

		cell = row.AddCell()
		cell.SetValue(rowData.VegasSpread)
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.SetValue(rowData.VegasOverUnder)
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.SetValue(rowData.VegasWinPercentage)
		cell.SetStyle(style)
	}

	actualHeaders := []string{"Team", "Spread", "O/U", "Outcome"}
	sheet.AddRow()
	actual := sheet.AddRow()
	for i, header := range actualHeaders {
		cell := actual.AddCell()
		style := xlsx.NewStyle()
		style.Border.Bottom = "thin"
		style.Font = *xlsx.NewFont(12, "Times New Roman")
		if i == 0 {
			style.Border.Right = "thick"
		}
		cell.Value = header
		cell.SetStyle(style)

	}

	for i, rowData := range goldenRows {
		style := xlsx.NewStyle()
		style.Font = *xlsx.NewFont(11, "Times New Roman")

		styleColumnLine := xlsx.NewStyle()
		styleColumnLine.Font = *xlsx.NewFont(11, "Times New Roman")
		styleColumnLine.Border.Right = "thick"
		if i%2 == 1 {
			style.Border.Bottom = "thin"
			styleColumnLine.Border.Bottom = "thin"
		}
		row := sheet.AddRow()

		cell := row.AddCell()
		cell.Value = rowData.Name
		cell.SetStyle(styleColumnLine)

		cell = row.AddCell()
		cell.SetValue("")
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.SetValue("")
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.SetValue("")
		cell.SetStyle(style)
	}

	// Save the XLSX file
	err = excelFile.Save("GoldenBoy.xlsx")
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
}
