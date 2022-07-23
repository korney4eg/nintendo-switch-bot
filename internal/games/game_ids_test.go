package games

import "testing"

func TestGetGameIDsDiff(t *testing.T) {
	type test struct {
		name    string
		first   []GameID
		second  []GameID
		diffLen int
	}
	var diff []GameID
	tests := []test{
		{
			name:    "single diff",
			first:   []GameID{"1", "2"},
			second:  []GameID{"5", "2"},
			diffLen: 1,
		},
		{
			name:    "different order",
			first:   []GameID{"1", "2", "3"},
			second:  []GameID{"1", "3", "2"},
			diffLen: 0,
		},
	}
	for _, tc := range tests {
		diff = GetGameIDsDiff(tc.first, tc.second)
		if len(diff) != tc.diffLen {
			t.Fatalf("[%s] Expected diff len: %d, got: %d", tc.name, tc.diffLen, len(diff))
		}

	}
}
