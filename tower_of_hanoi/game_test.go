package main

import "testing"

func TestNewGame(t *testing.T) {
	game, err := NewGame(3)
	if err != nil {
		t.Fatal(err)
	}
	if got := game.String(); got != "A: 3 2 1\nB: -\nC: -" {
		t.Fatalf("initial board = %q", got)
	}
}

func TestMoveRejectsInvalidMoves(t *testing.T) {
	game, _ := NewGame(2)
	for _, test := range []struct {
		name     string
		from, to int
	}{
		{"same peg", 0, 0},
		{"empty source", 1, 2},
	} {
		t.Run(test.name, func(t *testing.T) {
			if err := game.Move(test.from, test.to); err == nil {
				t.Fatal("Move unexpectedly succeeded")
			}
		})
	}

	if err := game.Move(0, 1); err != nil {
		t.Fatal(err)
	}
	if err := game.Move(0, 1); err == nil {
		t.Fatal("larger disk was placed on a smaller disk")
	}
}

func TestGameCanBeWon(t *testing.T) {
	game, _ := NewGame(2)
	for _, move := range [][2]int{{0, 1}, {0, 2}, {1, 2}} {
		if err := game.Move(move[0], move[1]); err != nil {
			t.Fatal(err)
		}
	}
	if !game.Won() {
		t.Fatal("game was not won")
	}
}

func TestParsePeg(t *testing.T) {
	for input, want := range map[string]int{"a": 0, " B ": 1, "C": 2} {
		if got, err := parsePeg(input); err != nil || got != want {
			t.Fatalf("parsePeg(%q) = %d, %v; want %d", input, got, err, want)
		}
	}
}
