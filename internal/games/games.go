package games

import (
	"encoding/json"
	"log"
	"os"
)

func LoadGamesFromFile(fileName string) (games Games, err error) {
	blob, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(blob, &games); err != nil {
		log.Fatal(err)
	}
	return games, err
}

func (games Games) GetAllTitles() (titles []string) {
	for _, game := range games {
		titles = append(titles, game.Title)
	}
  return titles
}
