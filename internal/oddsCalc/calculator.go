package oddsCalc

import (
	"math"
	"tommy2thicc/internal/schemas"
)

func CalcPotentialPayout(odd int, wagerAmount float64) float64 {
	return (calculateDecimalOdds(odd) * wagerAmount) - wagerAmount
}

func HouseProbs(odd1 int, odd2 int) (float64, float64) {
	houseProb1 := calculateHouseProbability(odd1)
	houseProb2 := calculateHouseProbability(odd2)

	realProb1 := houseProb1 / (houseProb1 + houseProb2)
	realProb2 := houseProb2 / (houseProb1 + houseProb2)

	return realProb1, realProb2
}

func ExpectedValFunc(prob float64, payout float64, wagerAmount float64) float64 {
	return ((prob * payout) - ((1 - prob) * wagerAmount))
}

func calculateHouseProbability(odds int) float64 {
	if odds < 0 {
		return float64((-1 * odds)) / float64((-1*odds)+100)
	} else {
		return float64(100) / float64(odds+100)
	}
}
func calculateDecimalOdds(odd int) float64 {
	if odd < 0 {
		return (float64(100) / float64(odd*-1)) + 1
	} else {
		return (float64(odd) / float64(100)) + 1
	}
}

func CalculateEloProb(hTElo float64, atElo float64) (float64, float64) {
	probHt := 1 / (1 + (math.Pow(10, ((atElo - (hTElo + 50)) / 400))))
	probAt := 1 - probHt
	return probHt, probAt
}

func CalculateKPPointDiff(homeTeam schemas.KPTeamStats, awayTeam schemas.KPTeamStats) (float64, float64) {
	hTpointDiff := ((homeTeam.Adjem - awayTeam.Adjem) * (awayTeam.AdjT + homeTeam.AdjT)) / 200
	aTpointDiff := ((awayTeam.Adjem - homeTeam.Adjem) * (awayTeam.AdjT + homeTeam.AdjT)) / 200

	return (hTpointDiff + 3.75), (aTpointDiff - 3.75)
}

func CalculateKPWinProb(marginDiff float64) float64 {
	return (.5) * (1 + math.Erf((marginDiff)/(11*math.Sqrt(2))))
}

func CalculateLog5KpWinProb(homeTeam schemas.KPTeamStats, awayTeam schemas.KPTeamStats) (float64, float64) {
	hTAdjO := (homeTeam.AdjO + (homeTeam.AdjO * .014))
	hTAdjD := (homeTeam.AdjD - (homeTeam.AdjD * .014))
	aTAdjO := (awayTeam.AdjO - (awayTeam.AdjO * .014))
	aTAdjD := (awayTeam.AdjD + (awayTeam.AdjD * .014))

	hTEw := (math.Pow(hTAdjO, 11.5) / ((math.Pow(hTAdjO, 11.5)) + math.Pow(hTAdjD, 11.5)))
	aTEw := (math.Pow(aTAdjO, 11.5) / ((math.Pow(aTAdjO, 11.5)) + math.Pow(aTAdjD, 11.5)))

	htPW := (hTEw - (hTEw * aTEw)) / ((hTEw + aTEw) - (2 * hTEw * aTEw))
	atPw := (aTEw - (aTEw * hTEw)) / ((aTEw + hTEw) - (2 * aTEw * hTEw))
	return htPW, atPw
}

func CalculateLog5KpPointMargin(homeTeam schemas.KPTeamStats, awayTeam schemas.KPTeamStats) (float64, float64) {
	hTAdjO := ((homeTeam.AdjO + (homeTeam.AdjO * .014)) / 101.5)
	hTAdjD := ((homeTeam.AdjD - (homeTeam.AdjD * .014)) / 101.5)
	aTAdjO := ((awayTeam.AdjO - (awayTeam.AdjO * .014)) / 101.5)
	aTAdjD := ((awayTeam.AdjD + (awayTeam.AdjD * .014)) / 101.5)

	htExOp := (hTAdjO * aTAdjD * 101.5)
	atExOp := (aTAdjO * hTAdjD * 101.5)

	return (htExOp * .697), (atExOp * .697)
}
