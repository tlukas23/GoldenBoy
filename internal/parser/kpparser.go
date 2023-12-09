package kpParser

import (
	"fmt"
	"html"
	"os"
	"regexp"
	"strconv"
	"strings"
	"tommy2thicc/internal/schemas"
)

func ParseKenPomHtml() (map[string]schemas.KPTeamStats, error) { // Read the HTML content from the text file
	fileContent, err := os.ReadFile("kpHtml.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	htmlContent := string(fileContent)

	// Extracting data from the HTML content
	rows := strings.Split(htmlContent, "</tr>")

	teamStats := make(map[string]schemas.KPTeamStats, 0)

	for _, row := range rows {
		if strings.Contains(row, "<td class=\"hard_left\">") {
			cells := strings.Split(row, "</td>")
			teamStat := schemas.KPTeamStats{
				TeamName: extractTextValue(cells[1]),
				Adjem:    extractFloatValue(cells[4]),
				AdjO:     extractFloatValue(cells[5]),
				AdjD:     extractFloatValue(cells[7]),
				AdjT:     extractFloatValue(cells[9]),
			}
			teamStats[teamStat.TeamName] = teamStat
		}
	}
	return teamStats, nil
}

// Function to extract text content without HTML tags
func extractTextValue(cell string) string {
	re := regexp.MustCompile(`<a.*?>(.*?)</a>`)
	matches := re.FindStringSubmatch(cell)
	if len(matches) > 1 {
		decoded := html.UnescapeString(matches[1])
		return decoded
	}
	return ""
}

func extractFloatValue(cell string) float64 {
	re := regexp.MustCompile(`[+-]?\d+(\.\d+)?`)
	numStr := re.FindString(cell)
	floatVal, _ := strconv.ParseFloat(numStr, 64)
	return floatVal
}
