package bot

import (
	"bufio"
	"fmt" //to print errors
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"drixevel.dev/bonk-bot/config" //importing our config package which we have created above

	"github.com/bwmarrin/discordgo" //discordgo package from the repo of bwmarrin .
	rcon "github.com/forewing/csgo-rcon"
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

	if m.Content == "Bonk!" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Bonk!")
	}

	//Makes sure the message has the prefix in front of it to register it as a command.
	if m.Content[0:1] != config.BotPrefix {
		return
	}

	//Strip the prefix from the message.
	m.Content = m.Content[1:]

	if m.Content == "help" {
		var helpMsg string

		helpMsg = "!help - Shows this message.\n"
		helpMsg += "!ping - pong\n"
		helpMsg += "!pong - ping\n"
		helpMsg += "!quote - Shows a random Scout quote.\n"
		helpMsg += "!connect/!join - Shows a link to connect to the server.\n"
		helpMsg += "!rcon - Allows you to send commands to the server. (Admin Only)\n"

		_, _ = s.ChannelMessageSend(m.ChannelID, helpMsg)
	} else if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	} else if m.Content == "pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "ping")
	} else if m.Content == "quote" {
		//Opens the quotes.txt file and reads it.
		//Each line contains a specific quote to show.
		file, err := os.Open("quotes.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		//Setup the file to be parsed through so we can pull a random line.
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
	} else if m.Content == "connect" || m.Content == "join" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Join the server: steam://connect/"+config.DNS)
	} else if strings.Contains(m.Content, "rcon") {

		if m.ChannelID != config.AdminChannel {
			_, _ = s.ChannelMessageSend(m.ChannelID, "You are not in the correct channel to use this command.")
			return
		}

		//Strip the 'RCON' part of the message, this can be done better but whatever.
		m.Content = m.Content[4:]

		_, _ = s.ChannelMessageSend(m.ChannelID, "Executing command: "+m.Content)

		//Connects via RCON to execute the command(s).
		c := rcon.New(config.IP, config.Password, time.Second*2)
		msg, err := c.Execute(m.Content)

		if err != nil {
			fmt.Println(err)
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, "Output: "+msg)
	}
}
