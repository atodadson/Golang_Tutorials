package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Getword() (randword string) {
	// Open the file
	file, err := os.Open("words.txt")
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

func allValuesGreen(m map[string]string) bool {
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
func ProcessGuess(guess string, word string) map[string]string {
	response := make(map[string]string)

	for index, char := range guess {
		if rune(word[index]) == char {
			response[string(char)] = "Green"
		} else if InWord(char, word) {
			response[string(char)] = "Yellow"
		} else {
			response[string(char)] = "Black"
		}
	}

	return response
}

func main() {
	randword := Getword()
	// fmt.Println("The random word is: ", randword)
	var guess string
	guesses_left := 7

	for i := 1; i < 7; i++ {
		guesses_left--
		fmt.Printf("[%v Guesses left] Guess the word: \n", guesses_left)
		fmt.Scan(&guess)
		response := ProcessGuess(guess, randword)

		// fmt.Println(response)
		fmt.Printf("%s:%s,  %s:%s,  %s:%s,  %s:%s, %s:%s \n",
			string(guess[0]), response[string(guess[0])],
			string(guess[1]), response[string(guess[1])],
			string(guess[2]), response[string(guess[2])],
			string(guess[3]), response[string(guess[3])],
			string(guess[4]), response[string(guess[4])],
		)
		if allValuesGreen(response) {
			fmt.Println("Yeeey, You got it right. That's a great guess")
			break
		}
	}
	fmt.Println("The word is: ", randword)
}
