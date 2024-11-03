package parser

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func HrParser() ([]GameInfo, error) {

	fileContent, err := os.ReadFile("oddsHtml.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	htmlSnippet := string(fileContent)
	// Regex patterns for extracting information
	awayTeamPattern := `<div class="show-for-medsmall">\d*\s*(.*?)</div>`
	homeTeamPattern := `<div class="show-for-medsmall">\d*\s*(.*?)</div>`
	spreadPattern := `<div class="column selection-line-value small-24 xxlarge-expand">\s*([+-]?\d+(\.\d+)?)\s*</div>`
	overUnderPattern := `<div class="row selection-container selection-container-vertical " data-cy="wager-button:Total Points (Over|Under)" data-tooltip-id="\d+-\d+"><div class="column selection-line-value small-24 xxlarge-expand">\s*[OU]\s*(\d+(\.\d+)?)\s*</div><div class="column selection-odds xxlarge-13 ">(-?\d+)</div></div>`
	moneyLinePattern := `<div class="row selection-container selection-container-vertical " data-cy="wager-button:To Win [AB]" data-tooltip-id="\d+-\d+"><div class="column selection-odds small-24 center-text">([+-]?\d+)</div></div>`

	games := make([]GameInfo, 0)
	// Extract information from each instance
	re := regexp.MustCompile(`<div class="row hr-market-view ".*?class="row infoBoxContainer"`)
	matches := re.FindAllString(htmlSnippet, -1)
	for _, match := range matches {
		//log.Println(match)
		// Extract away team name
		awayTeamRegex := regexp.MustCompile(awayTeamPattern)
		awayTeamMatches := awayTeamRegex.FindAllStringSubmatch(match, -1)
		awayTeamName := ""
		if len(awayTeamMatches) > 0 && len(awayTeamMatches[0]) >= 2 {
			re := regexp.MustCompile(`^#\d+\s+(.+)$`)
			matches := re.FindStringSubmatch(awayTeamMatches[0][1])
			if len(matches) > 1 {
				awayTeamName = matches[1]

			} else {
				awayTeamName = awayTeamMatches[0][1]
			}
		}

		// Extract home team name
		homeTeamRegex := regexp.MustCompile(homeTeamPattern)
		homeTeamMatches := homeTeamRegex.FindAllStringSubmatch(match, -1)
		homeTeamName := ""
		if len(homeTeamMatches) > 1 && len(homeTeamMatches[1]) >= 2 {
			re := regexp.MustCompile(`^#\d+\s+(.+)$`)
			matches := re.FindStringSubmatch(homeTeamMatches[1][1])
			if len(matches) > 1 {
				homeTeamName = matches[1]

			} else {
				homeTeamName = homeTeamMatches[1][1]
			}
		}

		game := GameInfo{
			AwayTeamName: TeamNameOptionValidator(awayTeamName),
			HomeTeamName: TeamNameOptionValidator(homeTeamName),
		}

		// Extract spread values
		spreadRegex := regexp.MustCompile(spreadPattern)
		spreadMatches := spreadRegex.FindAllStringSubmatch(match, -1)
		awayTeamSpread := 0.0
		homeTeamSpread := 0.0
		if len(spreadMatches) >= 2 && len(spreadMatches[0]) >= 2 && len(spreadMatches[1]) >= 2 {
			awayTeamSpread, _ = strconv.ParseFloat(spreadMatches[0][1], 64)
			homeTeamSpread, _ = strconv.ParseFloat(spreadMatches[1][1], 64)
		} else {
			log.Println("Bad spread match for the following game: ", awayTeamName, " vs ", homeTeamName)
			continue
		}
		game.AwayTeamSpread = awayTeamSpread
		game.HomeTeamSpread = homeTeamSpread

		// Extract over/under values
		overUnderRegex := regexp.MustCompile(overUnderPattern)
		overUnderMatches := overUnderRegex.FindAllStringSubmatch(match, -1)
		overUnderValue := 0.0
		if len(overUnderMatches) > 0 && len(overUnderMatches[0]) >= 3 {
			overUnderValue, _ = strconv.ParseFloat(overUnderMatches[0][2], 64)
		} else {
			log.Println("Bad O/U match for the following game: ", awayTeamName, " vs ", homeTeamName)
			continue
		}
		game.OverUnder = overUnderValue

		moneyLineRegex := regexp.MustCompile(moneyLinePattern)
		moneyLineMatches := moneyLineRegex.FindAllStringSubmatch(match, -1)
		var awayTeamMoneyLine float64
		var homeTeamMoneyLine float64
		if len(moneyLineMatches) >= 2 && len(moneyLineMatches[0]) >= 2 && len(moneyLineMatches[1]) >= 2 {
			awayTeamMoneyLine, _ = strconv.ParseFloat(moneyLineMatches[0][1], 64)
			homeTeamMoneyLine, _ = strconv.ParseFloat(moneyLineMatches[1][1], 64)
		}
		game.AwayTeamMl = float64(awayTeamMoneyLine)
		game.HomeTeamMl = float64(homeTeamMoneyLine)

		// Print the extracted information for each instance
		fmt.Println("Away Team Name:", game.AwayTeamName)
		fmt.Println("Home Team Name:", game.HomeTeamName)
		fmt.Printf("Away Team Spread: %.1f\n", awayTeamSpread)
		fmt.Printf("Home Team Spread: %.1f\n", homeTeamSpread)
		fmt.Printf("Over/Under Value: %.1f\n", overUnderValue)
		fmt.Println("Away Team Money Line: ", awayTeamMoneyLine)
		fmt.Println("Home Team Money Line: ", homeTeamMoneyLine)

		// Add a separator between instances
		fmt.Println(strings.Repeat("-", 40))
		games = append(games, game)
	}
	return games, nil
}
