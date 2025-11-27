package oddsCalc

import (
	"math/rand"
	"time"
	"tommy2thicc/internal/schemas"
)

func SimulateGame(teamA, teamB schemas.KPTeamStats, simulations int) float64 {
	teamAWins := 0

	seed := time.Now().UnixNano()
	rando := rand.New(rand.NewSource(seed))

	for i := 0; i < simulations; i++ {
		hTAdjO := ((teamA.AdjO + rando.NormFloat64()*9.04) / 101.5)
		hTAdjD := ((teamA.AdjD + rando.NormFloat64()*9.04) / 101.5)
		aTAdjO := ((teamB.AdjO + rando.NormFloat64()*9.04) / 101.5)
		aTAdjD := ((teamB.AdjD + rando.NormFloat64()*9.04) / 101.5)

		// htExOp := (hTAdjO * aTAdjD * 101.5 * .697)
		// atExOp := (aTAdjO * hTAdjD * 101.5 * .697)

		eTempo := (teamA.AdjT / 67.8) * (teamB.AdjT / 67.8) * 67.8
		htExOp := (hTAdjO * aTAdjD * 101.5 * (eTempo / 100))
		atExOp := (aTAdjO * hTAdjD * 101.5 * (eTempo / 100))

		// Simulate game outcome
		if htExOp > atExOp {
			teamAWins++
		}
	}

	// Return win probability
	return float64(teamAWins) / float64(simulations) * 100
}
