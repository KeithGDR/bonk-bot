package config

import (
	"encoding/json"
	"fmt"       //used to print errors majorly.
	"io/ioutil" //it will be used to help us read our config.json file.
)

var (
	Token        string //To store value of Token from config.json .
	BotPrefix    string // To store value of BotPrefix from config.json.
	IP           string // To store value of BotPrefix from config.json.
	Password     string // To store value of BotPrefix from config.json.
	AdminChannel string // To store value of BotPrefix from config.json.
	DNS          string // To store value of BotPrefix from config.json.

	config *configStruct //To store value extracted from config.json.
)

type configStruct struct {
	Token        string `json : "Token"`
	BotPrefix    string `json : "BotPrefix"`
	IP           string `json : "IP"`
	Password     string `json : "Password"`
	AdminChannel string `json : "AdminChannel"`
	DNS          string `json : "DNS"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	IP = config.IP
	Password = config.Password
	AdminChannel = config.AdminChannel
	DNS = config.DNS

	return nil
}
