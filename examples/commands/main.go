package main

import (
	"flag"
	"log"
    "os/exec"
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
	"github.com/whyrusleeping/hellabot/examples/commands/command"
	"github.com/whyrusleeping/hellabot/examples/commands/config"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

var irc_nick = "tekbot"
var bot_name = "gMiku"



// Flags for passing arguments to the program
var configFile = flag.String("config", "production.toml", "path to config file")

// core holds the command environment (bot connection and db)
var core *command.Core

// cmdList holds our command list, which tells the bot what to respond to.
var cmdList *command.List

// Main method
func main() {
	// Parse flags, this is needed for the flag package to work.
	// See https://godoc.org/flag
	flag.Parse()
	// Read the TOML Config
	conf := config.FromFile(*configFile)
	// Validate the config to see it's not missing anything vital.
	config.ValidateConfig(conf)

	// Setup our options anonymous function.. This gets called on the hbot.Bot object internally, applying the options inside.
	options := func(bot *hbot.Bot) {
		bot.SSL = conf.SSL
		if conf.ServerPassword != "" {
			bot.Password = conf.ServerPassword
		}
		bot.Channels = conf.Channels
	}
	// Create a new instance of hbot.Bot
	bot, err := hbot.NewBot(conf.Server, conf.Nick, options)
	if err != nil {
		log.Fatal(err)
	}
	// Setup the command environment
	core = &command.Core{bot, &conf}
	// Add the command trigger (this is what triggers all command handling)
	bot.AddTrigger(StringTrigger)
	bot.AddTrigger(JoinTrigger)
	// Set the default bot logger to stdout
	bot.Logger.SetHandler(log15.StdoutHandler)
	// Initialize the command list
	cmdList = &command.List{
		Prefix:   "!",
		Commands: make(map[string]command.Command),
	}

	// Start up bot (blocks until disconnect)
	bot.Run()
	log.Println("Bot shutting down.")
}

var JoinTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "JOIN"
	},
	func(bot *hbot.Bot, m *hbot.Message) bool {
		if m.From == irc_nick {
			// core.Bot.Reply(m, bot_name+": I'm "+bot_name+", ready to serve you")
			return false
		}
		return false
	},
}

// CommandTrigger passes all incoming messages to the commandList parser.
var StringTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG"
	},
	func(bot *hbot.Bot, m *hbot.Message) bool {
		if strings.Contains(m.From, irc_nick) {
			return false
		}
		if strings.Contains(strings.ToLower(m.Content), strings.ToLower(bot_name)) {
			core.Bot.Reply(m, bot_name+" is typing...")
			out, err := exec.Command("/bin/sh", "../../answer.sh", m.Content).Output()
			if err != nil {
				// log.Fatal(err)
			}
			core.Bot.Reply(m, string(out))
		}
		return false
	},
}
