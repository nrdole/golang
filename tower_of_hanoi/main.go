package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	disks := readDiskCount(scanner)
	game, err := NewGame(disks)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Move all disks from A to C. Enter moves like A C, or Q to quit.")
	for {
		fmt.Printf("\n%s\n", game)
		if game.Won() {
			fmt.Printf("You won in %d moves!\n", game.moves)
			return
		}

		fmt.Print("Move: ")
		if !scanner.Scan() || strings.EqualFold(strings.TrimSpace(scanner.Text()), "q") {
			fmt.Println("Thanks for playing!")
			return
		}
		parts := strings.Fields(scanner.Text())
		if len(parts) != 2 {
			fmt.Println("Enter a source and destination, such as A C.")
			continue
		}
		from, fromErr := parsePeg(parts[0])
		to, toErr := parsePeg(parts[1])
		if fromErr != nil || toErr != nil {
			fmt.Println("Enter pegs A, B, or C.")
			continue
		}
		if err := game.Move(from, to); err != nil {
			fmt.Println(err)
		}
	}
}

func readDiskCount(scanner *bufio.Scanner) int {
	fmt.Print("How many disks? (default 3): ")
	if !scanner.Scan() {
		return 3
	}
	value := strings.TrimSpace(scanner.Text())
	if value == "" {
		return 3
	}
	disks, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Invalid number; using 3 disks.")
		return 3
	}
	return disks
}
