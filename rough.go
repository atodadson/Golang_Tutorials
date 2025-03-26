package main

import "fmt"

func main() {
	// ANSI escape codes for different colors
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"
	aaa := "checkkls"

	// Print colored text
	fmt.Printf(red + "This is red text!" + "\n" + reset)
	fmt.Println(green + "This is green text!" + reset)
	fmt.Println(yellow + "This is yellow text!" + reset)
	fmt.Println("This is normal text!")
	fmt.Println(len(aaa))
}
