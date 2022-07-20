package games

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	// "log"
)

type GameStoreLocal struct {
	Games map[GameID]*Game
}

func (gameStore *GameStoreLocal) SaveToFile(fileName string) error {
	file, err := json.MarshalIndent(gameStore.Games, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, file, 0o644)
	return err
}

func (gameStore *GameStoreLocal) SaveIDsToFile(fileName string) error {
	games := []GameID{}
	for gameID := range gameStore.Games {
		games = append(games, gameID)
	}
	sort.Sort(ByID(games))
	gamesString := ""
	for _, game := range games {
		gamesString += fmt.Sprintf("%s\n", game)
	}

  err := ioutil.WriteFile(fileName, []byte(gamesString), 0o644)
	return err
}

func (gameStore *GameStoreLocal) LoadFromFile(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(file), &gameStore.Games)
	return err
}

func (gameStore *GameStoreLocal) AppendFromFile(fileName string) {
	var loadedGames map[string]*Game
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(file), &loadedGames)
	if err != nil {
		return
	}
	PrintIds(loadedGames)
	// log.Printf("Loaded games: %s", printIds(loadedGames))
	for _, loadedGame := range loadedGames {
		gameStore.AddGame(loadedGame)
	}
	return
}

func PrintIds(g map[string]*Game) string {
	ids := ""
	for game := range g {
		ids += fmt.Sprintf("%s\n", game)
	}
	return fmt.Sprintf("IDs : [\n%s]", ids)
}

func (gameStore *GameStoreLocal) AddGame(game *Game) {
	if len(gameStore.Games) == 0 {
		gameStore.Games = make(map[GameID]*Game)
	}
	if gameStore.Games[game.FsID] == nil {
		// log.Println(game.Title)
	}
	gameStore.Games[game.FsID] = game
}

func (gameStore *GameStoreLocal) HasGame(game *Game) bool {
	return gameStore.Games[game.FsID] != nil
}

func (gameStore *GameStoreLocal) FindGame(gameId GameID) (game *Game) {
	return gameStore.Games[gameId]
}
