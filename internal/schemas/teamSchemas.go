package schemas

// TeamStats represents the team statistics
type KPTeamStats struct {
	TeamName string
	Adjem    float64
	AdjO     float64
	AdjD     float64
	AdjT     float64
}

type CsvTeamOdds struct {
	TeamName  string
	Spread    float64
	OverUnder float64
	MoneyLine int
}

type GoldenCopyRow struct {
	Name                  string
	WinPercentageLog5     string
	PredictedPointsLog5   string
	KpWinPercentage       string
	KpSpread              string
	Log5PredictedTotal    string
	ExpectedMoneyLineLog5 string
	ExpectedMoneyLineKp   string
	VegasSpread           string
	VegasOverUnder        string
	VegasWinPercentage    string
}
