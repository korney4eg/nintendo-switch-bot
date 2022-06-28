package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/asdine/storm/v3"
	"github.com/korney4eg/nintendo-switch-bot/internal/games"
)

const (
	gamesPerPage = 48
	searchPath   = "select?q=*&fq=type%%3AGAME%%20AND%%20((playable_on_txt%%3A%%22HAC%%22))" +
		"%%20AND%%20sorting_title%%3A*%%20AND%%20*%%3A*&sort=deprioritise_b%%20asc" +
		"%%2C%%20popularity%%20asc&start=%d&rows=%d&wt=json" +
		"&bf=linear(ms(priority%%2CNOW%%2FHOUR)%%2C3.19e-11%%2C0)" +
		"&bq=!deprioritise_b%%3Atrue%%5E999"
)

func saveEachGame(game *games.Game, db *storm.DB) error {
	if err := db.Save(game); err != nil {
		return err
	}
	return nil
}

func downloadGames(host string) error {
	db, err := storm.Open("my.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	pageNum := 0
	gameNum := 0
	var nintendoResp *games.NintendoResponce
	for {
		// log.Println(pageNum)
		// log.Println(fmt.Sprintf(host+searchPath, pageNum*gamesPerPage, gamesPerPage))
		res, err := http.Get(fmt.Sprintf(host+searchPath, pageNum, gamesPerPage))
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// log.Printf("%s", body)
		if err = json.Unmarshal(body, &nintendoResp); err != nil {
			log.Println(err)
			break

		}
		// if pageNum == 1 {
		// 	break
		// }
		for _, game := range nintendoResp.Response.Docs {
			gameNum++
			log.Printf("%d/%d  Title: %s\n", gameNum, nintendoResp.Response.NumFound, game.Title)

			if err := saveEachGame(game, db); err != nil {
				return err
			}
		}
		if nintendoResp.Response.NumFound <= pageNum*gamesPerPage {
			break
		}
		pageNum++
	}
	return nil
}

func main() {
	log.Println("Starting")
	err := downloadGames("https://searching.nintendo-europe.com/ru/")
	if err != nil {
		log.Fatalln(err)
	}
}
