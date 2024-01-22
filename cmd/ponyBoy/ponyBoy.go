package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	Token := "MTE5ODg1MTU0NTI2MTkzNjY4MA.GFCH7a.qeHnMCcmkVdvWcEzZDeOkQEKak56vkyH_cmj1s"

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press Ctrl+C to exit.")
	// Wait here until CTRL+C or other term signal is received.
	select {
	case <-waitExit():
		return
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check if the message is from a bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!PrintMoney") {
		// Replace "path/to/your/file.txt" with the path to your file
		file, err := os.Open("GoldenBoy.xlsx")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		currentTime := time.Now()

		// Format the date as a string
		currentDateString := currentTime.Format("2006-01-02") // YYYY-MM-DD
		fileName := "GoldenBoy-" + currentDateString + ".xlsx"
		// Send the file to the channel
		s.ChannelFileSend(m.ChannelID, fileName, file)
	}
}

func waitExit() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	return c
}
