package main

import (
	"encoding/csv"
	"log"
	"os"
	"tommy2thicc/internal/kpParser"
	"tommy2thicc/internal/spreadsheet"
)

func main() {
	teamStats, err := kpParser.ParseKenPomHtml()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("NCAA_Mbb.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	goldenRows := spreadsheet.AnalyzeTheSheet(data[1:], teamStats)

	file, err := os.Create("GoldenBoy.csv")
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Team", "Expected ML Log5", "Expected ML Kp",
		"Predicted Points Log5", "Point Margin Kp", "Total Points Log5", "Hard Rock WP", "Log5 WP", "Kp WP"}
	if err := w.Write(headers); err != nil {
		log.Fatalln("error writing record to file", err)
	}
	for _, rowData := range goldenRows {
		row := []string{
			rowData.Name,
			rowData.ExpectedMoneyLineLog5,
			rowData.ExpectedMoneyLineKp,
			rowData.PredictedPointsLog5,
			rowData.KpPointMargin,
			rowData.Log5PredictedTotal,
			rowData.HardRockWinPercentage,
			rowData.WinPercentageLog5,
			rowData.KpWinPercentage,
		}
		log.Println(row)
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
