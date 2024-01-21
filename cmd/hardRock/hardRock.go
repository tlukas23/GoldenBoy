package main

import (
	"fmt"
	"log"
	"os"
	"time"
	kpParser "tommy2thicc/internal/parser"

	"github.com/tealeg/xlsx"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Game struct {
	AwayTeam   string
	OverUnder  string
	HomeTeam   string
	HomeSpread string
	AwaySpread string
}

func main() {
	filePath := "oddsHtml.txt"
	const (
		seleniumPath = "selenium/vendor/selenium-server.jar"
		chromeBinary = "chrome-linux64/chrome"
		chromeDriver = "chromedriver"
		port         = 4445 // Default port for Selenium WebDriver
	)

	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),         // Start an X frame buffer for the browser to run in.
		selenium.ChromeDriver(chromeDriver), // Specify the path to GeckoDriver in order to use Firefox.
	}

	chromeOpts := chrome.Capabilities{
		Path: chromeBinary,
	}

	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chromeOpts)
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("https://app.hardrock.bet/sport-leagues/basketball/691032891339243522"); err != nil {
		panic(err)
	}

	// Wait for a brief moment for the page to load (adjust this as needed)
	time.Sleep(5 * time.Second)

	// Get the page source (HTML content)
	htmlContent, err := wd.PageSource()
	if err != nil {
		fmt.Println("Failed to get page source:", err)
		return
	}

	err = os.WriteFile(filePath, []byte(htmlContent), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	gameOdds, err := kpParser.HrParser()
	if err != nil {
		log.Fatal(err)
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Games")
	if err != nil {
		fmt.Println("Error creating sheet:", err)
		return
	}

	// Add header row
	header := sheet.AddRow()
	header.AddCell().SetValue("Away Team")
	header.AddCell().SetValue("Away Spread")
	header.AddCell().SetValue("Over/Under")
	header.AddCell().SetValue("Away ML Odds")
	header.AddCell().SetValue("Home Team")
	header.AddCell().SetValue("Home Spread")
	header.AddCell().SetValue("Over/Under")
	header.AddCell().SetValue("Home ML Odds")

	for _, game := range gameOdds {
		row := sheet.AddRow()
		row.AddCell().SetValue(game.AwayTeamName)
		row.AddCell().SetFloat(game.AwayTeamSpread)
		row.AddCell().SetFloat(game.OverUnder)
		row.AddCell().SetValue(game.AwayTeamMl)
		row.AddCell().SetValue(game.HomeTeamName)
		row.AddCell().SetFloat(game.HomeTeamSpread)
		row.AddCell().SetFloat(game.OverUnder)
		row.AddCell().SetValue(game.HomeTeamMl)
	}

	// Save the Excel file
	err = file.Save("games.xlsx")
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
}
