package parser

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type GameFinal struct {
	Name   string
	Spread int
	Total  int
	Win    bool
}

func ParseHTML() []GameFinal {
	fileContent, err := os.ReadFile("espnHtml.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	html := string(fileContent)

	var teamFinal []GameFinal

	reW := regexp.MustCompile(`<tr class="winner"><td.*?team-name-link">(.*?)</a>.*?<td>(\d+)</td><td>(\d+)</td><td class="total">(\d+)</td>`)
	reL := regexp.MustCompile(`<tr class="loser"><td.*?team-name-link">(.*?)</a>.*?<td>(\d+)</td><td>(\d+)</td><td class="total">(\d+)</td>`)
	matchesW := reW.FindAllStringSubmatch(html, -1)
	matchesL := reL.FindAllStringSubmatch(html, -1)

	for i := range matchesW {
		log.Println(matchesW)
		Wname := strings.TrimSpace(matchesW[i][1])
		WtScore := parseInt(matchesW[i][4])
		Lname := strings.TrimSpace(matchesL[i][1])
		LtScore := parseInt(matchesL[i][4])

		if WtScore == 0 || LtScore == 0 {
			continue
		}

		teamW := GameFinal{Name: Wname, Spread: (LtScore - WtScore), Total: (WtScore + LtScore), Win: true}
		teamL := GameFinal{Name: Lname, Spread: (WtScore - LtScore), Total: (WtScore + LtScore), Win: false}
		teamFinal = append(teamFinal, teamW)
		teamFinal = append(teamFinal, teamL)
	}

	return teamFinal
}

func parseInt(s string) int {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		panic(err)
	}
	return n
}
