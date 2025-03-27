package main

import (
	"fmt"
	"os"
)

func main() {
	// ANSI escape codes for different colors
	// red := "\033[31m"
	// green := "\033[32m"
	// yellow := "\033[33m"
	// reset := "\033[0m"
	// aaa := "checkkls"

	// Print colored text
	// fmt.Printf(red + "This is red text!" + "\n" + reset)
	// fmt.Println(green + "This is green text!" + reset)
	// fmt.Println(yellow + "This is yellow text!" + reset)
	// fmt.Println("This is normal text!")
	// fmt.Println(len(aaa))

	// for i := 'A'; i <= 'Z'; i++ {
	//     fmt.Println(string(i))
	// }

	// for i := 0; i < 4; i++ {
	// 	fmt.Printf("Main loop: %v\n", i)
	// 	for j := 0; j < 4; j++ {
	// 		if j > 2 {
	// 			fmt.Printf("Will be breaking i: %v, j: %v\n", i, j)
	// 			break
	// 		} else {
	// 			fmt.Printf("Sub loop: %v\n\n", j)
	// 		}
	// 	}
	// }

	file, err := os.ReadFile("mywords.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		for index, word := range file {
			fmt.Println(index, string(word))
		}
		// fmt.Println(string(file))
	}

}
