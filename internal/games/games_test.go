package games

import "testing"

func TestGetDifference(t *testing.T) {
	marioGame2 := NewGame("2", "Mario2")
	marioGame3 := NewGame("2", "Mario3")
	marioGame3.PriceSortingF = 20.0
	differences := marioGame2.GetDifference(&marioGame3)
	if len(differences) < 1 {
		t.Error("There should be difference in Price, got no difference")
	}
	if differences[0].NewValue != float32(20) {
		t.Errorf("Difference in price should be 20.0, got %f\n", differences[0].NewValue)
	}
}
