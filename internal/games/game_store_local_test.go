package games

import (
	"testing"
)

func TestAddGame(t *testing.T) {
	gameStore := &GameStoreLocal{}
	marioGame := NewGame("1", "Mario1")
	gameStore.AddGame(&marioGame)
	luigiGame := NewGame("3", "Luigi")
	gameStore.AddGame(&luigiGame)

	if len(gameStore.Games) != 2 {
		t.Fatalf("There should be only 2 games, got %d\n", len(gameStore.Games))
	}
}

func TestHasGame(t *testing.T) {
	gameStore := &GameStoreLocal{}
	marioGame := NewGame("1", "Mario1")
	gameStore.AddGame(&marioGame)
	newGame := NewGame("6", "New Game")
	if gameStore.HasGame(&newGame) {
		t.Fatalf("New games shouldn't already be included into gameStore, but it is")
	}
	existingGame := NewGame("1", "Existing Game")
	if !gameStore.HasGame(&existingGame) {
		t.Error("Existing Game was not in list, but it should be")
	}
}

func TestFindGame(t *testing.T) {
	gameStore := &GameStoreLocal{}
	marioGame := NewGame("11", "Mario11")
	gameStore.AddGame(&marioGame)
	luigiGame := NewGame("13", "Luigy13")
	gameStore.AddGame(&luigiGame)
	existingGame := gameStore.FindGame("11")
	if existingGame == nil {
		t.Error("Couldn't find game with id 11, got nil")
	}
	nonExistingGame := gameStore.FindGame("12")
	if nonExistingGame != nil {
		t.Fatalf("found game with id %s, but should have nil\n", nonExistingGame.FsID)
	}
}

func TestGetChanges(t *testing.T) {
}
