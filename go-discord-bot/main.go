package main

import (
	"fmt"
	"go-discord-bot/commands"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	//init dot env for token
	env_err := godotenv.Load()
	if env_err != nil {
		fmt.Println("Error loading .env file")
	}

	//Init discord api object
	d, err := discordgo.New(os.Getenv("BOT_TOKEN"))

	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}

	//Add commands
	d.AddHandler(commands.HelloWorld)
	d.AddHandler(commands.Time)
	d.AddHandler(commands.RandomPokemon)

	//Read messages
	d.Identify.Intents = discordgo.IntentsGuildMessages

	//Open websoicket to listen
	err = d.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	d.Close()
}
