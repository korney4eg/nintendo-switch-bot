package games

import (
	"bufio"
	"fmt"
	"os"
)

type GameID string

type ByID []GameID

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i] < a[j] }

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadGameIDsFromFile(path string) ([]GameID, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var gameIds []GameID
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameIds = append(gameIds, GameID(scanner.Text()))
	}
	return gameIds, scanner.Err()
}

// writeLines writes the lines to the given file.
func WriteGameIDsToFile(gameIds []GameID, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range gameIds {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func GetGameIDsDiff(games1, games2 []GameID) []GameID {
	var diff []GameID

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for _, s1 := range games1 {
		found := false
		for _, s2 := range games2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if !found {
			diff = append(diff, s1)
		}
	}

	return diff
}
