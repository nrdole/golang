package main

import "fmt"

func main() {
	var num float64
	fmt.Print("Enter a floating point number: ")
	fmt.Scanf("%f", &num)
	fmt.Printf("The truncated value is: %d\n", int(num))
}