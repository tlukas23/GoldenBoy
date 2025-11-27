package spreadsheet

import (
	"log"
	"tommy2thicc/internal/oddsCalc"
	"tommy2thicc/internal/schemas"
)

func AnalyzeTheSheet(spreadsheet [][]string, teamsStats map[string]schemas.KPTeamStats) []schemas.GoldenCopyRow {
	goldenRows := make([]schemas.GoldenCopyRow, 0)
	for _, row := range spreadsheet {
		if row[0] == "" {
			break
		}
		htOdds, atOdds := ParseRowOdds(row)
		htStats, ok := teamsStats[htOdds.TeamName]
		if !ok {
			log.Println("Team name: ", htOdds.TeamName, " does not match anything in team stats list")
			continue
		}

		atStats, ok := teamsStats[atOdds.TeamName]
		if !ok {
			log.Println("Team name: ", atOdds.TeamName, " does not match anything in team stats list")
			continue
		}

		a := oddsCalc.SimulateGame(htStats, atStats, 10000)
		b := 100 - a
		log.Printf("\n%v: %.2f\n%v: %.2f\n\n", htOdds.TeamName, a, atOdds.TeamName, b)

		htHRWp, atHRWp := oddsCalc.HouseProbs(int(htOdds.MoneyLine), int(atOdds.MoneyLine))
		atLog5Spread, htLog5Spread := oddsCalc.CalculateLog5KpSpread(htStats, atStats)
		htLog5PredPts, atLog5PredPts := oddsCalc.BDDistSpread(htStats, atStats)
		totalPointsPred := htLog5PredPts + atLog5PredPts
		htKpPointDiff, atKpPointDiff := oddsCalc.CalculateKPPointDiff(htStats, atStats)

		kpHtPointSpread := htKpPointDiff * -1
		kpAtPointSpread := atKpPointDiff * -1

		htRow := schemas.GoldenCopyRow{
			Name:                htOdds.TeamName,
			VegasWinPercentage:  htHRWp * 100,
			PredictedPointsLog5: htLog5PredPts,
			KpSpread:            kpHtPointSpread,
			Log5PredictedTotal:  totalPointsPred,
			VegasSpread:         htOdds.Spread,
			VegasOverUnder:      htOdds.OverUnder,
			Log5Spread:          htLog5Spread,
		}

		atRow := schemas.GoldenCopyRow{
			Name:                atOdds.TeamName,
			VegasWinPercentage:  atHRWp * 100,
			PredictedPointsLog5: atLog5PredPts,
			KpSpread:            kpAtPointSpread,
			Log5PredictedTotal:  totalPointsPred,
			VegasSpread:         atOdds.Spread,
			VegasOverUnder:      atOdds.OverUnder,
			Log5Spread:          atLog5Spread,
		}

		goldenRows = append(goldenRows, atRow)
		goldenRows = append(goldenRows, htRow)
	}

	return goldenRows
}
