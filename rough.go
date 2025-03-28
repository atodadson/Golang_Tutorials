package main

import (
	"bufio"
	"fmt"
	"os"
)

func Get_Dict() (dictionary []string) {
	// Open the first file
	file, err := os.Open("new_words.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read all words into a slice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dictionary = append(dictionary, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	return
}

func WordIn_Dict(guess string, dictionary []string) bool {
	for _, word := range dictionary {
		if guess == word {
			return true
		}
	}
	return false
}

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
	// REGEX pattern for matching words
	// var pattern = "^[A-Z]+$"
	// var reg, _ = regexp.Compile(pattern)

	array := Get_Dict()
	fmt.Printf("%v %v", array, len(array))

	indict := WordIn_Dict("pumma", array)
	fmt.Println(indict)

	for i := 1; i < 1000; i++ {
		fmt.Printf("%v %v\n", i, string(i))
	}

}
