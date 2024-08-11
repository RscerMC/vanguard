package main

import (
	"fmt"
	"os"
	"os/signal"
	"vanguard/bot"
	"vanguard/common"
)

func init() {
	err := common.Init()
	if err != nil {
		panic(err)
	}
}

func main() {
	err := bot.Session.Open()
	if err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	bot.Session.Close()
	fmt.Println("Bot is shutting down.")

}
