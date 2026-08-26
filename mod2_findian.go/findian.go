package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var formatString string
	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')
	formatString = strings.ToLower(strings.TrimSpace(inputString))

	if strings.HasPrefix(formatString, "i") && strings.Contains(formatString, "a") && strings.HasSuffix(formatString, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
