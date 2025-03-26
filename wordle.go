package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Getword() (randword string, wordlist []string) {
	// Open the file
	file, err := os.Open("mywords.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read all words into a slice
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Generate a random index
	rand.Seed(time.Now().UnixNano()) // Seed the random generator
	randomIndex := rand.Intn(len(words))

	randword = words[randomIndex]
	wordlist = words
	return
}

func InWord(char rune, word string) bool {
	for _, character := range word {
		if char == character {
			return true
		}
	}
	return false
}

func allValuesGreen(m [5]string) bool {
	for _, value := range m {
		if value != "Green" {
			return false
		}
	}
	return true
}

// func ProcessGuess (guess string, word string) (response array){
// 	for index, char := range guess {
// 			if word[index] == char {
// 				response[char] = "Green"
// 			} else if Inword(char, word) {
// 				response[char] = "Yellow"
// 			} else {
// 				response[char] = "Black"
// 			}
// 	}
// }

// ProcessGuess processes the guess and returns a response
func ProcessGuess(guess string, word string) [5]string {
	response := [5]string{}

	for index, char := range guess {
		if rune(word[index]) == char {
			response[index] = "Green"
		} else if InWord(char, word) {
			response[index] = "Yellow"
		} else {
			response[index] = "Red"
		}
	}
	return response
}

func GetColour(colour_name string) (colour string) {
	// ANSI escape codes for different colors
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	if colour_name == "Green" {
		colour = green
	} else if colour_name == "Yellow" {
		colour = yellow
	} else if colour_name == "Red" {
		colour = red
	} else {
		panic(fmt.Sprintf("GetColour Func received unexpected input: %s", colour))
	}
	return
}

func main() {
	reset := "\033[0m"
	randword, _ := Getword()
	randword = strings.ToLower(randword)
	// fmt.Println("The random word is: ", randword)
	fmt.Printf("---------Welcome to Dadson's wordle game-------------\n------Guess the randomly selected 5 letter word------\n")
	var guess string
	guesses_left := 7

	for i := 1; i < 7; i++ {
		guesses_left--
		fmt.Printf("[%v Guesses left] Guess the word: \n", guesses_left)
		fmt.Scan(&guess)
		response := ProcessGuess(guess, randword)

		// fmt.Println(response)
		fmt.Printf(GetColour(response[0]) + string(guess[0]) + " " + reset)
		fmt.Printf(GetColour(response[1]) + string(guess[1]) + " " + reset)
		fmt.Printf(GetColour(response[2]) + string(guess[2]) + " " + reset)
		fmt.Printf(GetColour(response[3]) + string(guess[3]) + " " + reset)
		fmt.Printf(GetColour(response[4]) + string(guess[4]) + " " + "\n" + reset)
		if allValuesGreen(response) {
			fmt.Println("Yeeey, You got it right. That's a great guess")
			break
		}
	}
	fmt.Println("The word is: ", randword)
}
