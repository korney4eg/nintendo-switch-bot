package games

import "testing"

func TestPagination(t *testing.T) {
	marioGame1 := NewGame("1", "Mario1")
	marioGame2 := NewGame("2", "Mario2")
	marioGame3 := NewGame("3", "Mario3")
	marioGame4 := NewGame("4", "Mario4")
	marioGame5 := NewGame("5", "Mario5")
	marioGames := []*Game{
		&marioGame1, &marioGame2,
		&marioGame3, &marioGame4,
		&marioGame5,
	}
	paginatedMario := Paginate(marioGames, 2)
	t.Logf("marioGames = %+v", marioGames)
	if len(paginatedMario.Pages) != 3 {
		t.Logf("First element of GamesList is '%+v'", paginatedMario.Pages[0])
		t.Fatalf("paginatedMario expected 3 pages, got %d\n full list = '%+v'", len(paginatedMario.Pages), paginatedMario.Pages)
	}
  if len(paginatedMario.Pages[2]) != 1 {
    t.Errorf("Last page of Mario Games size supposed to be 1, but got %d\n", len(paginatedMario.Pages[2]))

  }
}
