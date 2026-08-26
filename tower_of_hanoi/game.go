package main

import (
	"fmt"
	"strings"
)

type Game struct {
	pegs  [3][]int
	total int
	moves int
}

func NewGame(disks int) (*Game, error) {
	if disks < 1 {
		return nil, fmt.Errorf("the game needs at least one disk")
	}

	game := &Game{total: disks}
	for disk := disks; disk > 0; disk-- {
		game.pegs[0] = append(game.pegs[0], disk)
	}
	return game, nil
}

func (g *Game) Move(from, to int) error {
	if from < 0 || from >= len(g.pegs) || to < 0 || to >= len(g.pegs) {
		return fmt.Errorf("peg must be A, B, or C")
	}
	if from == to {
		return fmt.Errorf("choose two different pegs")
	}
	if len(g.pegs[from]) == 0 {
		return fmt.Errorf("the source peg is empty")
	}

	disk := g.pegs[from][len(g.pegs[from])-1]
	if len(g.pegs[to]) > 0 && g.pegs[to][len(g.pegs[to])-1] < disk {
		return fmt.Errorf("you cannot place a larger disk on a smaller one")
	}

	g.pegs[from] = g.pegs[from][:len(g.pegs[from])-1]
	g.pegs[to] = append(g.pegs[to], disk)
	g.moves++
	return nil
}

func (g *Game) Won() bool {
	return len(g.pegs[2]) == g.total
}

func (g *Game) String() string {
	var lines []string
	for i, name := range []string{"A", "B", "C"} {
		contents := "-"
		if len(g.pegs[i]) > 0 {
			values := make([]string, len(g.pegs[i]))
			for j, disk := range g.pegs[i] {
				values[j] = fmt.Sprint(disk)
			}
			contents = strings.Join(values, " ")
		}
		lines = append(lines, fmt.Sprintf("%s: %s", name, contents))
	}
	return strings.Join(lines, "\n")
}

func parsePeg(value string) (int, error) {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "A":
		return 0, nil
	case "B":
		return 1, nil
	case "C":
		return 2, nil
	default:
		return 0, fmt.Errorf("peg must be A, B, or C")
	}
}
