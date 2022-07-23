/*
Copyright © 2022 Aliaksei Karneyeu korneevayu@gmail.com

*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/korney4eg/nintendo-switch-bot/internal/games"
	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls games data from nintendo e-shop",
	Run: func(cmd *cobra.Command, args []string) {
		err := downloadGames(host)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var (
	startPageNum int
	gamesPerPage int
	dataFile     string
	idsFile      string
)

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().IntVarP(&startPageNum, "page", "p", 1, "Page number to pull data from. Data from that single page would be pulled")
	pullCmd.Flags().IntVarP(&gamesPerPage, "rows", "r", 100, "Rows per page to download")
	pullCmd.Flags().StringVarP(&dataFile, "data", "d", "data/games.json", "path to json file with games data")
	pullCmd.Flags().StringVarP(&idsFile, "ids-file", "i", "data/ids", "path to file with games ids")
}

const (
	host       = "https://searching.nintendo-europe.com/ru/"
	searchPath = "select?q=*&fq=type%%3AGAME%%20AND%%20((playable_on_txt%%3A%%22HAC%%22))" +
		"%%20AND%%20sorting_title%%3A*%%20AND%%20*%%3A*&sort=deprioritise_b%%20asc" +
		"%%2C%%20popularity%%20asc&start=%d&rows=%d&wt=json" +
		"&bf=linear(ms(priority%%2CNOW%%2FHOUR)%%2C3.19e-11%%2C0)" +
		"&bq=!deprioritise_b%%3Atrue%%5E999"
)

var UnsupportedStatusCode = errors.New("Unsupported status code")

func GetEShopPage(host string, pageNum, gamesPerPage int) ([]byte, error) {
	res, err := http.Get(fmt.Sprintf(host+searchPath, pageNum, gamesPerPage))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil, UnsupportedStatusCode
	}
	return body, nil
}

func downloadGames(host string) error {
	// TODO:
	// 1. Page Iterator
	// 2. Page Reader
	// 3. Page processor

	GameStore := &games.GameStoreLocal{}
	pageNum := startPageNum
	checkedGames := []games.GameID{}
	GameStore.LoadFromFile(dataFile)
	var nintendoResp *games.NintendoResponce
	for {
		page, err := GetEShopPage(host, pageNum*gamesPerPage, gamesPerPage)
		if err != nil {
			log.Fatalf("http query error: %v\n", err)
		}
		if err = json.Unmarshal(page, &nintendoResp); err != nil {
			log.Println(err)
			break

		}
		for _, game := range nintendoResp.Response.Docs {
			checkedGames = append(checkedGames, game.FsID)
			if !GameStore.HasGame(game) {
				log.Printf("[INFO] %d. Adding (%s) %s\n", len(checkedGames), game.FsID, game.Title)
				GameStore.AddGame(game)
			}
		}
		if nintendoResp.Response.NumFound <= pageNum*gamesPerPage {
			break
		}
		if startPageNum != 1 {
			break
		}
		pageNum++
	}
	log.Printf("Checked %d games\n", len(checkedGames))
	if err := GameStore.SaveToFile(dataFile); err != nil {
		return err
	}
	firstRun := false
	previouslyCheckedGames, err := games.ReadGameIDsFromFile(idsFile)
	if err != nil {
		log.Printf("[ERROR] couldn't open file: %s\n", idsFile)
		firstRun = true
	}
	if !firstRun {
		gameIDsDifference := games.GetGameIDsDiff(checkedGames, previouslyCheckedGames)
		log.Println("Added game ids:")
		for _, gameID := range gameIDsDifference {
			log.Printf("- %s\n", gameID)
		}
		gameIDsDifference = games.GetGameIDsDiff(previouslyCheckedGames, checkedGames)
		log.Println("Removed game ids:")
		for _, gameID := range gameIDsDifference {
			log.Printf("- %s\n", gameID)
		}
	}
	if err := games.WriteGameIDsToFile(checkedGames, idsFile); err != nil {
		log.Printf("[ERROR] couldn't save file: %s\n", idsFile)
		return err
	}
	return nil
}
