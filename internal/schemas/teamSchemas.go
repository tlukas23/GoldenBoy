package schemas

// TeamStats represents the team statistics
type KPTeamStats struct {
	TeamName string
	Adjem    float64
	AdjO     float64
	AdjD     float64
	AdjT     float64
	Luck     float64
}

type CsvTeamOdds struct {
	TeamName  string
	Spread    float64
	OverUnder float64
	MoneyLine int
}

type GoldenCopyRow struct {
	Name                string
	PredictedPointsLog5 float64
	KpSpread            float64
	Log5Spread          float64
	Log5PredictedTotal  float64
	VegasSpread         float64
	VegasOverUnder      float64
	VegasWinPercentage  float64
}
