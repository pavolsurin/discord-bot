package main

import (
	"github.com/pavolsurin/discord-bot/pkg/bot"
)

func main() {
	bot.Bot()
	<-make(chan struct{})
}
