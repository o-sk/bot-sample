package main

import (
	"log"
	"os"

	"github.com/sbstjn/hanu"
)

var SlackApiToken string

func init() {
	SlackApiToken = os.Getenv("SLACK_API_TOKEN")
}

func main() {
	bot, err := hanu.New(SlackApiToken)

	if err != nil {
		log.Fatal(err)
	}

	bot.Command("hi", func(conv hanu.ConversationInterface) {
		conv.Reply("hi")
	})

	bot.Listen()
}
