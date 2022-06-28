package telegram

import (
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"

	"github.com/asdine/storm/v3"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/korney4eg/nintendo-switch-bot/internal/commands"
	"github.com/korney4eg/nintendo-switch-bot/internal/games"
)

var logger *zap.Logger

func InitLogger() {
	logger, _ = zap.NewDevelopment()
}

func helpScreen() string {
	messageText := "Available commands:\n/help - to see this screen"
	messageText += "\n/all - to show all games"
	messageText += "\n/categories - to show categories"
	return messageText
}

func startScreen() string {
	messageText := "Welcome to Nintendo eShop bot"
	return messageText + "\n" + helpScreen()
}

func GameScreen(game games.Game) string {
	messageText := game.Title
	return messageText + "\n" + helpScreen()
}

func getGame(game *games.Game) string {
	messageText := fmt.Sprintf("name: *%s*\n", game.Title)
	messageText += fmt.Sprintf("https://www.nintendo.ru/-%s", game.URL)
	return messageText
}

func SendMessage(chatId int64, message string) {
}

func MainLoop() {
	fmt.Println("Starting main loop")
	InitLogger()
	db, err := storm.Open("my.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log := logger.Sugar()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Fatal("Error loading telegram client", zap.Error(err))
	}
	log.Info("Loading games data")
	gamesList := &games.GamesList{}
	gamesData, err := games.LoadGamesFromFile("all_games.json")
	gamesList.Games = gamesData
	if err != nil {
		logger.Fatal("Error loading games", zap.Error(err))
	}
	log.Info("Games data successfully loaded")

	log.Infof("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// var messages []*tgbotapi.MessageConfig
	// paginatedGames := games.Paginate(gameData, 5)
	// page := 0
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		if update.Message.IsCommand() {
			// messages = messages[:0]
			// 	if update.Message.Command() == "all" {
			// 		for _, game := range paginatedGames.Pages[page] {
			// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			// 			msg.Text = views.ShowGame(game)
			// 			msg.ParseMode = tgbotapi.ModeMarkdown
			// 			msg.DisableWebPagePreview = true
			// 			messages = append(messages, &msg)
			// 		}
			// 	}
			// 	if update.Message.Command() == "next" {
			// 		page++
			// 		for _, game := range paginatedGames.Pages[page] {
			// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			// 			msg.Text = views.ShowGame(game)
			// 			msg.ParseMode = tgbotapi.ModeMarkdown
			// 			// msg.DisableWebPagePreview = true
			// 			messages = append(messages, &msg)
			// 		}
			// 	}
			// AllCommands := &commands.CommandsList{Commands: []*commands.Command{&commands.StartCommand, &commands.HelpCommand, commands.GetAllGamesCommand()}}
			// command := AllCommands.FindCommand(update.Message.Command(), db)
			// log.Infof("got command: %s", command.View.Text)
			messages := commands.Execute(update.Message.Chat.ID, update.Message.Command(), db)

			for _, message := range messages {
				if msg, err := bot.Send(message); err != nil {
					log.Error("Error sending telegram message", zap.Error(err), " msg ", fmt.Sprintf("%+v", msg))
				}
			}
		}

	}
}
