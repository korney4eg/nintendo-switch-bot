package games

import (
	"encoding/json"
	"log"
	"os"
)

func (game *Game) GetDifference(otherGame *Game) (differences []*Difference) {
	if game.PriceSortingF != otherGame.PriceSortingF {
		differences = append(differences, &Difference{
			Field:    "Price",
			OldValue: game.PriceSortingF,
			NewValue: otherGame.PriceSortingF,
		})
	}
	return differences
}

func LoadGamesFromFile(fileName string) (games []*Game, err error) {
	blob, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(blob, &games); err != nil {
		log.Fatal(err)
	}
	return games, err
}

func NewGame(id GameID, title string) (game Game) {
	game.FsID = id
	game.Title = title
	return game
}

type ByID []GameID

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i] < a[j] }
