package kpParser

import (
	"fmt"
	"html"
	"log"
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
	HomeTeamMl     float64
	AwayTeamMl     float64
}

func FDParser() ([]GameInfo, error) {

	fileContent, err := os.ReadFile("oddsHtml.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	htmlText := string(fileContent)

	aTeamNameRegex := regexp.MustCompile(`"awayTeam"\s*:\s*{\s*"@type"\s*:\s*"SportsTeam",\s*"name"\s*:\s*"([^"]*)"\s*}`)
	homeTeamRegex := regexp.MustCompile(`"homeTeam"\s*:\s*{\s*"@type"\s*:\s*"SportsTeam",\s*"name"\s*:\s*"([^"]*)"\s*}`)

	aTeamMatches := aTeamNameRegex.FindAllStringSubmatch(htmlText, -1)
	hTeamMatches := homeTeamRegex.FindAllStringSubmatch(htmlText, -1)
	games := make([]GameInfo, 0)

	for i := 0; i < len(aTeamMatches); i++ {
		// Populate game info
		game := GameInfo{
			AwayTeamName: aTeamMatches[i][1],
			HomeTeamName: hTeamMatches[i][1],
		}

		aTregexPattern := fmt.Sprintf(`aria-label="%s, (-?\d+\.?\d*),`, html.EscapeString(game.AwayTeamName))
		hTregexPattern := fmt.Sprintf(`aria-label="%s, (-?\d+\.?\d*),`, html.EscapeString(game.HomeTeamName))
		reAt := regexp.MustCompile(aTregexPattern)
		aTmatches := reAt.FindAllStringSubmatch(htmlText, -1)
		reHt := regexp.MustCompile(hTregexPattern)
		hTmatches := reHt.FindAllStringSubmatch(htmlText, -1)

		if len(aTmatches) == 0 || len(hTmatches) == 0 {
			log.Println("No spread available for team game", game.AwayTeamName, "@", game.HomeTeamName)
			continue
		}

		game.AwayTeamSpread, err = strconv.ParseFloat(aTmatches[0][1], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing Away Team Spread value: %w", err)
		}

		game.HomeTeamSpread, err = strconv.ParseFloat(hTmatches[0][1], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing Home Team Spread value: %w", err)
		}

		overUnderpattern := fmt.Sprintf(`%s.*?Over (\d+\.\d+).*`, html.EscapeString(game.AwayTeamName))
		reOU := regexp.MustCompile(overUnderpattern)
		ouMatches := reOU.FindAllStringSubmatch(htmlText, -1)

		if len(ouMatches) == 0 {
			log.Println("No over under available for team game", game.AwayTeamName, "@", game.HomeTeamName)
			continue
		}
		overUnder, err := strconv.ParseFloat(ouMatches[0][1], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing OverUnder value: %w", err)
		}
		game.OverUnder = overUnder

		htMlregex := fmt.Sprintf(`aria-label="%s, ([+-]?\d+(\.\d+)?) Odds"`, html.EscapeString(game.HomeTeamName))
		atMlregex := fmt.Sprintf(`aria-label="%s, ([+-]?\d+(\.\d+)?) Odds"`, html.EscapeString(game.AwayTeamName))
		reAtMl := regexp.MustCompile(atMlregex)
		aTMlmatches := reAtMl.FindAllStringSubmatch(htmlText, -1)
		reHtMl := regexp.MustCompile(htMlregex)
		hTMlmatches := reHtMl.FindAllStringSubmatch(htmlText, -1)

		if len(hTMlmatches) == 0 || len(aTMlmatches) == 0 {
			log.Println("No money line available for team game", game.AwayTeamName, "@", game.HomeTeamName)
			continue
		}

		game.AwayTeamMl, err = strconv.ParseFloat(aTMlmatches[0][1], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing Away Team Spread value: %w", err)
		}

		game.HomeTeamMl, err = strconv.ParseFloat(hTMlmatches[0][1], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing Home Team Spread value: %w", err)
		}

		game.AwayTeamName = strings.Replace(game.AwayTeamName, "State", "St.", 1)
		game.HomeTeamName = strings.Replace(game.HomeTeamName, "State", "St.", 1)

		game.AwayTeamName = TeamNameOptionValidator(game.AwayTeamName)
		game.HomeTeamName = TeamNameOptionValidator(game.HomeTeamName)

		//log.Println(game)
		games = append(games, game)
	}

	return games, nil
}

func TeamNameOptionValidator(name string) string {
	if name == "Miami" {
		return "Miami FL"
	}
	if name == "Ole Miss" {
		return "Mississippi"
	}
	if name == "Nebraska Cornhuskers" {
		return "Nebraska"
	}
	if name == "Pennsylvania" {
		return "Penn"
	}
	if name == "CSU Northridge" {
		return "Cal St. Northridge"
	}
	if name == "Long Island University" {
		return "LIU"
	}
	if name == "California Baptist" {
		return "Cal Baptist"
	}
	if name == "UIC" {
		return "Illinois Chicago"
	}
	if name == "UL Monroe" {
		return "Louisiana Monroe"
	}
	if name == "Arkansas-Pine Bluff" {
		return "Arkansas Pine Bluff"
	}
	if name == "SE Louisiana" {
		return "Southeastern Louisiana"
	}
	if name == "NC St." {
		return "N.C. State"
	}
	if name == "UT Martin" {
		return "Tennessee Martin"
	}
	if name == "Florida International" {
		return "FIU"
	}
	if name == "McNeese" {
		return "McNeese St."
	}
	if name == "WV Mountaineers" {
		return "West Virginia"
	}
	if name == "Central Conn. St." {
		return "Central Connecticut"
	}
	if name == "Southern Methodist" {
		return "SMU"
	}
	if name == "Charlotte 49ers" {
		return "Charlotte"
	}
	if name == "St Josephs" {
		return "Saint Joseph's"
	}
	if name == "Citadel" {
		return "The Citadel"
	}
	if name == "Gardner-Webb" {
		return "Gardner Webb"
	}
	if name == "Nicholls" {
		return "Nicholls St."
	}
	if name == "CSU Bakersfield" {
		return "Cal St. Bakersfield"
	}

	return name
}
