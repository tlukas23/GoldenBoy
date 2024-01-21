package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
	kpParser "tommy2thicc/internal/parser"

	"github.com/tealeg/xlsx"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
	var input string
	flag.StringVar(&input, "date", "", "HTML string to parse")
	flag.Parse()

	filePath := "espnHtml.txt"
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
		Args: []string{
			"--headless", // Run in headless mode
			"--disable-gpu",
		},
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

	if input != "" {
		if err := wd.Get("https://www.cbssports.com/college-basketball/scoreboard/FBS/" + input); err != nil {
			panic(err)
		}
	} else {
		// Navigate to the simple playground interface.
		if err := wd.Get("https://www.cbssports.com/college-basketball/scoreboard/"); err != nil {
			panic(err)
		}
	}

	// Wait for a brief moment for the page to load (adjust this as needed)
	time.Sleep(1 * time.Second)

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

	teams := kpParser.ParseHTML()
	fmt.Println(teams)

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Games")
	if err != nil {
		fmt.Println("Error creating sheet:", err)
		return
	}

	actualHeaders := []string{"Team", "Spread", "O/U", "Outcome"}
	sheet.AddRow()
	actual := sheet.AddRow()
	for i, header := range actualHeaders {
		cell := actual.AddCell()
		style := xlsx.NewStyle()
		style.Border.Bottom = "thin"
		style.Font = *xlsx.NewFont(12, "Times New Roman")
		if i == 0 {
			style.Border.Right = "thick"
		}
		cell.Value = header
		cell.SetStyle(style)

	}

	for i, rowData := range teams {
		style := xlsx.NewStyle()
		style.Font = *xlsx.NewFont(11, "Times New Roman")

		styleColumnLine := xlsx.NewStyle()
		styleColumnLine.Font = *xlsx.NewFont(11, "Times New Roman")
		styleColumnLine.Border.Right = "thick"
		if i%2 == 1 {
			style.Border.Bottom = "thin"
			styleColumnLine.Border.Bottom = "thin"
		}
		row := sheet.AddRow()

		cell := row.AddCell()
		cell.Value = rowData.Name
		cell.SetStyle(styleColumnLine)

		cell = row.AddCell()
		cell.Value = strconv.Itoa(rowData.Spread)
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.Value = strconv.Itoa(rowData.Total)
		cell.SetStyle(style)

		cell = row.AddCell()
		if rowData.Win {
			cell.Value = "W"
		} else {
			cell.Value = "L"
		}
		cell.SetStyle(style)
	}

	// Save the Excel file
	err = file.Save("outcomes.xlsx")
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
}
