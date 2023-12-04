package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {

	filePath := "kpHtml.txt"
	// Start a Selenium WebDriver server instance (you should have the server running)
	const (
		seleniumPath = "/home/tlukas/selenium/vendor/selenium-server.jar"
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
	if err := wd.Get("https://kenpom.com/"); err != nil {
		panic(err)
	}

	// Wait for a brief moment for the page to load (adjust this as needed)
	time.Sleep(5 * time.Second)

	// Get the page source (HTML content)
	html, err := wd.PageSource()
	if err != nil {
		fmt.Println("Failed to get page source:", err)
		return
	}

	err = os.WriteFile(filePath, []byte(html), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
