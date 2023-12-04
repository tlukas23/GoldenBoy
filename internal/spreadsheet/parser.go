package spreadsheet

import (
	"strconv"
	"tommy2thicc/internal/schemas"
)

func ParseRowOdds(rowData []string) (schemas.CsvTeamOdds, schemas.CsvTeamOdds) {
	ht := schemas.CsvTeamOdds{}
	at := schemas.CsvTeamOdds{}
	for x := range rowData {
		switch x {
		case 1:
			odd1, _ := strconv.Atoi(rowData[1])
			odd2, _ := strconv.Atoi(rowData[5])

			at.Spread = odd1
			ht.Spread = odd2
		case 2:
			odd1, _ := strconv.Atoi(rowData[2])
			odd2, _ := strconv.Atoi(rowData[6])
			at.OverUnder = odd1
			ht.OverUnder = odd2

		case 3:
			odd1, _ := strconv.Atoi(rowData[3])
			odd2, _ := strconv.Atoi(rowData[7])

			at.MoneyLine = odd1
			ht.MoneyLine = odd2
		}
		ht.TeamName = rowData[4]
		at.TeamName = rowData[0]
	}

	return ht, at
}
