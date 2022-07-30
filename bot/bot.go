package bot

import (
	"bufio"
	"fmt" //to print errors
	"log"
	"math/rand"
	"os"
	"time"

	"drixevel.dev/bonk-bot/config" //importing our config package which we have created above

	"github.com/bwmarrin/discordgo" //discordgo package from the repo of bwmarrin .
)

var BotId string
var goBot *discordgo.Session

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot has been created.")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself.
	if m.Author.ID == BotId {
		return
	}

	//Makes sure the message has the prefix in front of it to register it as a command.
	if m.Content[0:1] != config.BotPrefix {
		return
	}

	//Strip the prefix from the message.
	m.Content = m.Content[1:]

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	} else if m.Content == "pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "ping")
	} else if m.Content == "Bonk!" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Bonk!")
	} else if m.Content == "quote" {
		file, err := os.Open("quotes.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()

			roll := randgenerator.Intn(lineNum)

			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, pick)
	}
}
