package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// The slice starts empty, with room for three integers.
	numbers := make([]int, 0, 3)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter an integer (or X to quit): ")
		if !scanner.Scan() {
			return
		}

		input := strings.TrimSpace(scanner.Text())
		if strings.EqualFold(input, "X") {
			return
		}

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter an integer or X.")
			continue
		}

		numbers = append(numbers, number)
		sort.Ints(numbers)
		fmt.Println(numbers)
	}
}
