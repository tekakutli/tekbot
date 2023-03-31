package command

import (
	"strings"

	"github.com/whyrusleeping/hellabot/examples/commands/config"

	hbot "github.com/whyrusleeping/hellabot"
)

// Core holds the environment passed to each command handler
type Core struct {
	Bot    *hbot.Bot
	Config *config.Config
}

// Command represents a single command the bot will handle
type Command struct {
	Name        string
	Description string
	Usage       string
	Run         Func
}

// Func represents the Go function that will be executed when a command triggers
type Func func(m *hbot.Message, args []string)

// List holds the command list and prefix
type List struct {
	Prefix   string
	Commands map[string]Command
}

// AddCommand adds a command to the bots internal list
func (cl *List) AddCommand(c Command) {
	cl.Commands[c.Name] = c
}

// Process handles incoming messages and looks for incoming messages that start with the command prefix. Commands are triggered if they exist
func (cl *List) Process(bot *hbot.Bot, m *hbot.Message) {
	// Is the first character our command prefix?
	if m.Content[:1] == cl.Prefix {
		parts := strings.Fields(m.Content[1:])
		commandstring := parts[0]
		cmd, ok := cl.Commands[commandstring]
		if !ok {
			return
		}
		// looks good, get the quote and reply with the result
		bot.Logger.Debug("action", "start processing",
			"args", parts,
			"full text", m.Content)
		go func(m *hbot.Message) {
			bot.Logger.Debug("action", "executing",
				"full text", m.Content)
			if len(parts) > 1 {
				cmd.Run(m, parts[1:])
			} else {
				cmd.Run(m, []string{})
			}
		}(m)
	}
}
