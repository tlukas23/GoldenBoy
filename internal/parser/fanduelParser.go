package kpParser

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type GameInfo struct {
	AwayTeamName   string
	HomeTeamName   string
	OverUnder      float64
	AwayTeamSpread float64
	HomeTeamSpread float64
}

func FDParser() ([]GameInfo, error) {

	fileContent, err := os.ReadFile("oddsHtml.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	htmlText := string(fileContent)

	teamNameRegex := regexp.MustCompile(`<span aria-label="([^"]+)" role="text" class="s t ht[^"]*">([^<]+)</span>`)
	overUnderPattern := regexp.MustCompile(`Over (\d+\.\d+)`)
	spreadRegex := regexp.MustCompile(`>([-+]?\d+\.\d+)<`)

	teamMatches := teamNameRegex.FindAllStringSubmatch(htmlText, -1)
	overUnderMatches := overUnderPattern.FindAllStringSubmatch(htmlText, -1)
	spreadMatches := spreadRegex.FindAllString(htmlText, -1)
	numGames := len(teamMatches) / 2
	games := make([]GameInfo, numGames)

	// Extract game data
	for i := 0; i < numGames; i++ {
		// Populate game info
		game := GameInfo{
			AwayTeamName: teamMatches[i*2][1],
			HomeTeamName: teamMatches[i*2+1][1],
		}

		overUnder, err := strconv.ParseFloat(overUnderMatches[i][1], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing OverUnder value: %w", err)
		}
		game.OverUnder = overUnder
		game.AwayTeamSpread, err = strconv.ParseFloat(strings.Trim(spreadMatches[i*2], "><"), 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing Away Team Spread value: %w", err)
		}

		game.HomeTeamSpread, err = strconv.ParseFloat(strings.Trim(spreadMatches[i*2+1], "><"), 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing Home Team Spread value: %w", err)
		}
		games[i] = game
	}

	return games, nil
}
