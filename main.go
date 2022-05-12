package main

import (
	"log"

	"github.com/korney4eg/nintendo-switch-bot/internal/games"
)

func main() {
	games, err := games.LoadGamesFromFile("all_games_f.json")
	if err != nil {
		log.Fatalln(err)
	}
  for _, title := range (games.GetAllTitles()){
    log.Println(title)

  }
}
