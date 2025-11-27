package main

import (
	"fmt"
	"log"
	kpParser "tommy2thicc/internal/parser"
	"tommy2thicc/internal/spreadsheet"

	"github.com/tealeg/xlsx"
)

const (
	StyleLightGreen = iota
	Yellow
	StyleWhite
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

	headers := []string{"Team",
		"Predicted Points Log5", "Spread Kp", "Spread Log5", "Total Points Log5", "", "Team", "Vegas Spread", "Vegas O/U", "Vegas WP"}

	headerRow := sheet.AddRow()
	for i, header := range headers {
		style := xlsx.NewStyle()
		style.Border.Bottom = "thin"
		style.Font = *xlsx.NewFont(12, "Times New Roman")
		if i == 0 || i == 4 || i == 6 {
			style.Border.Right = "thick"
		}
		cell := headerRow.AddCell()
		cell.Value = header
		cell.SetStyle(style)
		// Set the column width to accommodate the content
		if i != 5 {
			sheet.SetColWidth(i, i, float64(13)) // Set the width for the first column (0-based index)
		} else {
			sheet.SetColWidth(i, i, float64(4)) // Set the width for the first column (0-based index)
		}
	}

	for i, rowData := range goldenRows {
		row := sheet.AddRow()

		cellName := row.AddCell()
		cellName.Value = rowData.Name
		style := styleCell(i, true, StyleWhite)
		cellName.SetStyle(style)

		ppLog := row.AddCell()
		ppLog.SetValue(rowData.PredictedPointsLog5)
		style = styleCell(i, false, StyleWhite)
		ppLog.SetStyle(style)

		kpSpread := row.AddCell()
		kpSpread.SetValue(rowData.KpSpread)
		if rowData.KpSpread < 0 && rowData.VegasSpread > 0 || rowData.KpSpread > 0 && rowData.VegasSpread < 0 {
			style = styleCell(i, false, Yellow)
		} else if rowData.KpSpread > rowData.VegasSpread+3.5 || rowData.KpSpread < rowData.VegasSpread-3.5 {
			style = styleCell(i, false, StyleLightGreen)
		} else {
			style = styleCell(i, false, StyleWhite)
		}
		kpSpread.SetStyle(style)

		log5Spread := row.AddCell()
		log5Spread.SetValue(rowData.Log5Spread)
		if rowData.Log5Spread < 0 && rowData.VegasSpread > 0 || rowData.Log5Spread > 0 && rowData.VegasSpread < 0 {
			style = styleCell(i, false, Yellow)
		} else if rowData.Log5Spread > rowData.VegasSpread+5.5 || rowData.Log5Spread < rowData.VegasSpread-5.5 {
			style = styleCell(i, false, StyleLightGreen)
		} else {
			style = styleCell(i, false, StyleWhite)
		}
		log5Spread.SetStyle(style)

		log5Predi := row.AddCell()
		log5Predi.SetValue(rowData.Log5PredictedTotal)
		if rowData.Log5PredictedTotal > rowData.VegasOverUnder+6 || rowData.Log5PredictedTotal < rowData.VegasOverUnder-6 {
			style = styleCell(i, true, StyleLightGreen)
		} else {
			style = styleCell(i, true, StyleWhite)
		}
		log5Predi.SetStyle(style)

		cell := row.AddCell()
		cell.Value = ""

		cell = row.AddCell()
		cell.Value = rowData.Name
		style = styleCell(i, true, StyleWhite)
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.SetValue(rowData.VegasSpread)
		style = styleCell(i, false, StyleWhite)
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.SetValue(rowData.VegasOverUnder)
		style = styleCell(i, false, StyleWhite)
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.SetValue(rowData.VegasWinPercentage)
		style = styleCell(i, false, StyleWhite)
		cell.SetStyle(style)
	}

	// Save the XLSX file
	err = excelFile.Save("GoldenBoy.xlsx")
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
}

func styleCell(i int, colLine bool, colorCell int) *xlsx.Style {
	style := xlsx.NewStyle()
	style.Font = *xlsx.NewFont(11, "Times New Roman")
	if colLine {
		style.Border.Right = "thick"
	}

	if i%2 == 1 {
		style.Border.Bottom = "thin"
	}

	switch colorCell {
	case StyleLightGreen:
		style.Fill = *xlsx.FillGreen
	case Yellow:
		style.Fill = *xlsx.FillGreen
		style.Fill.FgColor = "FFFFD966"
	}
	return style
}
