package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const RLSite = map[string]{
	OfficialSite: "https://www.rocketleague.com/en",
	RLCSUpdates:  "https://en.wikipedia.org/wiki/Rocket_League_Championship_Series#Seasons", // need to find proper link for updates here but we can work with this up until web scraping
}

func main() {
	// discord session && bot token
	dg, err := discordgo.New("Bot YOUR_BOT_TOKEN_HERE")
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// message creation
	dg.AddHandler(messageCreate)

	// discord connection
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection to Discord:", err)
		return
	}

	// run this sucker
	fmt.Println("Bot is now running. Press CTRL+C to end bot connection.")
	<-make(chan struct{})
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//  message handling logic here using m.Content
	// ex: m.Content = "message"
	m.Content = "Hello, I am RLB, I am a Bot. If you do not wish to see my messages, make sure you disable channel notifications at: #rocket-league > Notifications > Only @mentions "
}
