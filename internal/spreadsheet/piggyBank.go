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

		htHRWp, atHRWp := oddsCalc.HouseProbs(int(htOdds.MoneyLine), int(atOdds.MoneyLine))
		htLog5Wp, atLog5Wp := oddsCalc.CalculateLog5KpWinProb(htStats, atStats)

		atLog5Spread, htLog5Spread := oddsCalc.CalculateLog5KpSpread(htStats, atStats)
		htLog5PredPts, atLog5PredPts := oddsCalc.BDDistSpread(htStats, atStats)
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
			VegasWinPercentage:    htHRWp * 100,
			WinPercentageLog5:     htLog5Wp,
			PredictedPointsLog5:   htLog5PredPts,
			KpWinPercentage:       htKpWp,
			KpSpread:              kpHtPointSpread,
			ExpectedMoneyLineLog5: htLog5MlExpect,
			ExpectedMoneyLineKp:   htKpMlExpect,
			Log5PredictedTotal:    totalPointsPred,
			VegasSpread:           htOdds.Spread,
			VegasOverUnder:        htOdds.OverUnder,
			Log5Spread:            htLog5Spread,
		}

		atRow := schemas.GoldenCopyRow{
			Name:                  atOdds.TeamName,
			VegasWinPercentage:    atHRWp * 100,
			WinPercentageLog5:     atLog5Wp,
			PredictedPointsLog5:   atLog5PredPts,
			KpWinPercentage:       atKpWp,
			KpSpread:              kpAtPointSpread,
			ExpectedMoneyLineLog5: atLog5MlExpect,
			ExpectedMoneyLineKp:   atKpMlExpect,
			Log5PredictedTotal:    totalPointsPred,
			VegasSpread:           atOdds.Spread,
			VegasOverUnder:        atOdds.OverUnder,
			Log5Spread:            atLog5Spread,
		}

		goldenRows = append(goldenRows, atRow)
		goldenRows = append(goldenRows, htRow)
	}

	return goldenRows
}
