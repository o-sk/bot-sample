package main

import (
	"log"
	"math/rand"
	"os"

	"github.com/o-sk/gis"
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
	bot.Command("image <word>", func(conv hanu.ConversationInterface) {
		str, _ := conv.String("word")
		images, err := gis.Search(str)
		if err != nil {
			conv.Reply("Oops")
			return
		}
		image := images[rand.Intn(len(images))]
		if image.Source == "" {
			conv.Reply("Sorry, I can't find %s", str)
			return
		}
		conv.Reply("%s", image.Source)
	})

	bot.Listen()
}
