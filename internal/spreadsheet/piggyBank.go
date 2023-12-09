package spreadsheet

import (
	"log"
	"strconv"
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

		htHRWp, atHRWp := oddsCalc.HouseProbs(int(htOdds.MoneyLine), int(atOdds.MoneyLine))
		htLog5Wp, atLog5Wp := oddsCalc.CalculateLog5KpWinProb(htStats, atStats)
		htLog5PredPts, atLog5PredPts := oddsCalc.CalculateLog5KpSpread(htStats, atStats)
		totalPointsPred := htLog5PredPts + atLog5PredPts
		htKpPointDiff, atKpPointDiff := oddsCalc.CalculateKPPointDiff(htStats, atStats)
		htKpWp := oddsCalc.CalculateKPWinProb(htKpPointDiff)
		atKpWp := oddsCalc.CalculateKPWinProb(atKpPointDiff)

		htPayoutMl := oddsCalc.CalcPotentialPayout(htOdds.MoneyLine, 5)
		atPayoutMl := oddsCalc.CalcPotentialPayout(atOdds.MoneyLine, 5)

		htLog5MlExpect := oddsCalc.ExpectedValFunc(htLog5Wp, htPayoutMl, 5)
		atLog5MlExpect := oddsCalc.ExpectedValFunc(atLog5Wp, atPayoutMl, 5)
		htKpMlExpect := oddsCalc.ExpectedValFunc(htKpWp, htPayoutMl, 5)
		atKpMlExpect := oddsCalc.ExpectedValFunc(atKpWp, atPayoutMl, 5)

		kpHtPointSpread := htKpPointDiff * -1
		kpAtPointSpread := atKpPointDiff * -1

		htRow := schemas.GoldenCopyRow{
			Name:                  htOdds.TeamName,
			VegasWinPercentage:    strconv.FormatFloat(htHRWp, 'f', -1, 64),
			WinPercentageLog5:     strconv.FormatFloat(htLog5Wp, 'f', -1, 64),
			PredictedPointsLog5:   strconv.FormatFloat(htLog5PredPts, 'f', -1, 64),
			KpWinPercentage:       strconv.FormatFloat(htKpWp, 'f', -1, 64),
			KpSpread:              strconv.FormatFloat(kpHtPointSpread, 'f', -1, 64),
			ExpectedMoneyLineLog5: strconv.FormatFloat(htLog5MlExpect, 'f', -1, 64),
			ExpectedMoneyLineKp:   strconv.FormatFloat(htKpMlExpect, 'f', -1, 64),
			Log5PredictedTotal:    strconv.FormatFloat(totalPointsPred, 'f', -1, 64),
			VegasSpread:           strconv.FormatFloat(htOdds.Spread, 'f', -1, 64),
			VegasOverUnder:        strconv.FormatFloat(htOdds.OverUnder, 'f', -1, 64),
		}

		atRow := schemas.GoldenCopyRow{
			Name:                  atOdds.TeamName,
			VegasWinPercentage:    strconv.FormatFloat(atHRWp, 'f', -1, 64),
			WinPercentageLog5:     strconv.FormatFloat(atLog5Wp, 'f', -1, 64),
			PredictedPointsLog5:   strconv.FormatFloat(atLog5PredPts, 'f', -1, 64),
			KpWinPercentage:       strconv.FormatFloat(atKpWp, 'f', -1, 64),
			KpSpread:              strconv.FormatFloat(kpAtPointSpread, 'f', -1, 64),
			ExpectedMoneyLineLog5: strconv.FormatFloat(atLog5MlExpect, 'f', -1, 64),
			ExpectedMoneyLineKp:   strconv.FormatFloat(atKpMlExpect, 'f', -1, 64),
			Log5PredictedTotal:    strconv.FormatFloat(totalPointsPred, 'f', -1, 64),
			VegasSpread:           strconv.FormatFloat(atOdds.Spread, 'f', -1, 64),
			VegasOverUnder:        strconv.FormatFloat(atOdds.OverUnder, 'f', -1, 64),
		}

		goldenRows = append(goldenRows, atRow)
		goldenRows = append(goldenRows, htRow)
	}

	return goldenRows
}
