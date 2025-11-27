package oddsCalc

import (
	"math"
	"tommy2thicc/internal/schemas"
)

func HouseProbs(odd1 int, odd2 int) (float64, float64) {
	houseProb1 := calculateHouseProbability(odd1)
	houseProb2 := calculateHouseProbability(odd2)

	realProb1 := houseProb1 / (houseProb1 + houseProb2)
	realProb2 := houseProb2 / (houseProb1 + houseProb2)

	return realProb1, realProb2
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

func CalculateKPPointDiff(homeTeam schemas.KPTeamStats, awayTeam schemas.KPTeamStats) (float64, float64) {
	hTpointDiff := ((homeTeam.Adjem - awayTeam.Adjem) * (awayTeam.AdjT + homeTeam.AdjT)) / 200
	aTpointDiff := ((awayTeam.Adjem - homeTeam.Adjem) * (awayTeam.AdjT + homeTeam.AdjT)) / 200

	return hTpointDiff, aTpointDiff
}

func CalculateLog5KpWinProb(homeTeam schemas.KPTeamStats, awayTeam schemas.KPTeamStats) (float64, float64) {
	hTAdjO := (homeTeam.AdjO + (homeTeam.AdjO * .016))
	hTAdjD := (homeTeam.AdjD - (homeTeam.AdjD * .016))
	aTAdjO := (awayTeam.AdjO - (awayTeam.AdjO * .016))
	aTAdjD := (awayTeam.AdjD + (awayTeam.AdjD * .016))
	// hTAdjO := ((homeTeam.AdjO) / 101.5)
	// hTAdjD := ((homeTeam.AdjD) / 101.5)
	// aTAdjO := ((awayTeam.AdjO) / 101.5)
	// aTAdjD := ((awayTeam.AdjD) / 101.5)

	hTEw := (math.Pow(hTAdjO, 11.5) / ((math.Pow(hTAdjO, 11.5)) + math.Pow(hTAdjD, 11.5)))
	aTEw := (math.Pow(aTAdjO, 11.5) / ((math.Pow(aTAdjO, 11.5)) + math.Pow(aTAdjD, 11.5)))

	htPW := (hTEw - (hTEw * aTEw)) / ((hTEw + aTEw) - (2 * hTEw * aTEw))
	atPw := (aTEw - (aTEw * hTEw)) / ((aTEw + hTEw) - (2 * aTEw * hTEw))
	return htPW, atPw
}

func CalculateLog5KpSpread(homeTeam schemas.KPTeamStats, awayTeam schemas.KPTeamStats) (float64, float64) {
	hTAdjO := ((homeTeam.AdjO + (homeTeam.AdjO * .016)) / 101.5)
	hTAdjD := ((homeTeam.AdjD - (homeTeam.AdjD * .016)) / 101.5)
	aTAdjO := ((awayTeam.AdjO - (awayTeam.AdjO * .016)) / 101.5)
	aTAdjD := ((awayTeam.AdjD + (awayTeam.AdjD * .016)) / 101.5)

	// hTAdjO := ((homeTeam.AdjO) / 101.5)
	// hTAdjD := ((homeTeam.AdjD) / 101.5)
	// aTAdjO := ((awayTeam.AdjO) / 101.5)
	// aTAdjD := ((awayTeam.AdjD) / 101.5)

	eTempo := (homeTeam.AdjT / 68.1) * (awayTeam.AdjT / 68.1) * 68.1
	htExOp := (hTAdjO * aTAdjD * 101.5 * (eTempo / 100))
	atExOp := (aTAdjO * hTAdjD * 101.5 * (eTempo / 100))

	return (htExOp - atExOp), (atExOp - htExOp)
}

func BDDistSpread(homeTeam schemas.KPTeamStats, awayTeam schemas.KPTeamStats) (float64, float64) {
	// hTAdjO := (homeTeam.AdjO + (homeTeam.AdjO * .014))
	// hTAdjD := (homeTeam.AdjD - (homeTeam.AdjD * .014))
	// aTAdjO := (awayTeam.AdjO - (awayTeam.AdjO * .014))
	// aTAdjD := (awayTeam.AdjD + (awayTeam.AdjD * .014))
	hTAdjO := (homeTeam.AdjO)
	hTAdjD := (homeTeam.AdjD)
	aTAdjO := (awayTeam.AdjO)
	aTAdjD := (awayTeam.AdjD)

	poss := (homeTeam.AdjT * awayTeam.AdjT) / 72.4

	htPPP := (hTAdjO * aTAdjD) / 101.5
	atPPP := (aTAdjO * hTAdjD) / 101.5

	htPoints := (htPPP * poss) / 101.5
	atPoints := (atPPP * poss) / 101.5

	return htPoints, atPoints
}
