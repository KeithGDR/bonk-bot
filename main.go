package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"drixevel.dev/bonk-bot/bot"
	"drixevel.dev/bonk-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}
