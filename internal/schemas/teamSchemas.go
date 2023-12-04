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
	Spread    int
	OverUnder int
	MoneyLine int
}

type GoldenCopyRow struct {
	Name                  string
	WinPercentageLog5     string
	PredictedPointsLog5   string
	KpWinPercentage       string
	KpPointMargin         string
	Log5PredictedTotal    string
	HardRockWinPercentage string
	ExpectedMoneyLineLog5 string
	ExpectedMoneyLineKp   string
}
