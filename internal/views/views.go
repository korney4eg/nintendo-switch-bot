package views

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/korney4eg/nintendo-switch-bot/internal/games"
)

const nintendoRuName = "https://www.nintendo.ru/"

type (
	View struct {
		Text     string
		NextView *View
	}
)


var StartScreen = View{
	Text:     "Welcome to Nintendo eShop bot",
	NextView: &HelpScreen,
}

func GetUnknownCommandScreen(command string) *View {
	return &View{
		Text:     fmt.Sprintf("Unkown command: `%s`\n", command),
		NextView: &HelpScreen,
	}
}

var UnknownScreen = View{
	Text:     "Unkown command",
	NextView: &HelpScreen,
}

var AllGamesScreen = View{
	Text: "Here you get all games",
}

// var CategoriesScreen = View{
// 	Text:     games.GetCategories,
// }

func GetGameScreen(game *games.Game) (view *View) {
	view.Text = fmt.Sprintf("**[%s](%s/%s)**\nPrice: %f/%f", game.Title, nintendoRuName, game.URL, game.PriceLowestF, game.PriceRegularF)
	return view
}

func ShowGame(game *games.Game, message *tgbotapi.MessageConfig) {
	// return fmt.Sprintf("**[%s](%s/%s)**\nPrice: %d/%d", game.Title, nintendoRuName, game.URL, game.PriceLowestF, game.PriceRegularF)
}

var HelpScreen = View{
	Text: `Available commands:
/help - to see this screen
/all - to show all games(will display 5 games)
/categories - to show categories
/next - to display next set of games`,
}

func (view View) Show(game *games.Game) {}
