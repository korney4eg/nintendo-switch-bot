package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/asdine/storm/v3"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/korney4eg/nintendo-switch-bot/internal/views"
)

type (
	Command struct {
		Name              string
		View              *views.View
		AvailableCommands []*Command
		CommandFounded    func(string) bool
		Process           func(*storm.DB, string) error
	}
	CommandsList struct {
		Commands []*Command
	}
)

func (command *Command) SetText(text string) {
	if command.View == nil {
		command.View = &views.View{}
	}
	command.View.Text = text
}

var StartCommand = Command{
	View: &views.StartScreen,
	CommandFounded: func(cmd string) bool {
		return cmd == "start"
	},
	Process: func(db *storm.DB, cmd string) error {
		return nil
	},
}

var HelpCommand = Command{
	View: &views.HelpScreen,
	CommandFounded: func(cmd string) bool {
		return cmd == "help"
	},
	Process: func(db *storm.DB, cmd string) error {
		return nil
	},
}

var GetGameCommand = Command{
	Name: "game_*",
	Process: func(db *storm.DB, cmd string) error {
		return nil
	},
}

func GetUnknowCommand(command string) *Command {
	return &Command{
		View: views.GetUnknownCommandScreen(command),
		Process: func(db *storm.DB, cmd string) error {
			return nil
		},
	}
}

func GetFollowCommand(cmdLine string) *Command {
	command := &Command{
		View: &views.View{
			Text: "empty",
		},
		CommandFounded: func(cmd string) bool {
			return strings.HasPrefix(cmd, "follow_game_")
		},
	}
	command.Process = func(db *storm.DB, cmd string) error {
		log.Printf("Processing command `%s`", cmd)
		command.SetText(fmt.Sprintf("Processing `%s` ...", cmd))
		return nil
	}
	return command
}

func GetAllGamesCommand() *Command {
	allGamesView := &views.AllGamesScreen
	return &Command{
		View: allGamesView,
		CommandFounded: func(cmd string) bool {
			return cmd == "all"
		},
		Process: func(db *storm.DB, cmd string) error {
			return nil
		},
	}
}

// func GetGame(game *games.Game) {
// 	return &Command{}
// }

func Execute(chatId int64, cmd string) (messages []*tgbotapi.MessageConfig) {
	// view := &views.UnknownScreen
	// command := FindCommand(cmd)
	// log.Printf("[INFO] Executing command `%s`\n", cmd)
	// if command.View != nil {
	// 	view = command.View
	// }
	// command.Process(db, cmd)
	// log.Println("Processed")

	// for {
	// 	msg := tgbotapi.NewMessage(chatId, "")
	// 	msg.Text = view.Text
	// 	msg.ParseMode = tgbotapi.ModeMarkdown
	// 	msg.DisableWebPagePreview = true
	// 	messages = append(messages, &msg)
	// 	if view.NextView != nil {
	// 		view = view.NextView
	// 	} else {
	// 		break
	// 	}
	// }
	return messages
}

func FindCommand(providedCommand string) *Command {
	AllCommands := &CommandsList{
		Commands: []*Command{
			&StartCommand,
			&HelpCommand,
			GetFollowCommand(providedCommand),
		},
	} // AddCommand  = &CommandsList{Commands: []*Command{&StartCommand, &HelpCommand}}
	for _, command := range AllCommands.Commands {
		if command.CommandFounded(providedCommand) {
			return command
		}
	}
	return GetUnknowCommand(providedCommand)
}
