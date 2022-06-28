package games

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func insertStudent(db *sql.DB, code string, name string, program string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, program)
	if err != nil {
		log.Fatalln(err.Error())
	}
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

// func (games GamesList) GetAllTitles() (titles []string) {
// 	for _, game := range games.Games {
// 		titles = append(titles, game.Title)
// 	}
// 	return titles
// }

func NewGame(id, title string) (game Game) {
	game.FsID = id
	game.Title = title
	return game
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetCategories(games []*Game) (categories []string) {
	for _, game := range games {
		for _, category := range game.GameCategoriesTxt {
			if !stringInSlice(category, categories) {
				categories = append(categories, category)
			}
		}
	}
	return categories
}

func Paginate(games []*Game, size int) (paginatedGames Paginated) {
	gamesOnPage := []*Game{}
	for gameNum, game := range games {
		if len(game.DatesReleasedDts) > 0 && game.DatesReleasedDts[0].Before(time.Now()) {
			gamesOnPage = append(gamesOnPage, game)
		}
		if gameNum == len(games)-1 && len(gamesOnPage) < size {
			paginatedGames.Pages = append(paginatedGames.Pages, gamesOnPage)
			return paginatedGames
		}
		if gameNum <= len(games)-1 && len(gamesOnPage) == size {
			paginatedGames.Pages = append(paginatedGames.Pages, gamesOnPage)
			gamesOnPage = []*Game{}
		}
	}
	return paginatedGames
}
