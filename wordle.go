package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"regexp"
)

// ANSI escape codes for different colors
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var reset = "\033[0m"

// REGEX pattern for matching words
var pattern = "^[A-Z]+$"
var reg, _ = regexp.Compile(pattern)

// Keyboard display
func CreateKeyboard() (keyboard map[string]string) {
    keyboard = make(map[string]string, 6)
    for i := 'A'; i <= 'Z'; i++ {
        keyboard[string(i)] = reset
    }
    return keyboard
}

func ShowKeyboard(keyboard map[string]string) {
    row1 := []string{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P"}
    row2 := []string{"A", "S", "D", "F", "G", "H", "J", "K", "L"}
    row3 := []string{"Z", "X", "C", "V", "B", "N", "M"}
    for _, letter := range row1 {
        fmt.Printf(keyboard[letter] + letter + " " + reset )
    }
    fmt.Println("")
    for _, letter := range row2 {
        fmt.Printf(keyboard[letter] + letter + " " + reset )
    }
    fmt.Println("")
    for _, letter := range row3 {
        fmt.Printf(keyboard[letter] + letter + " " + reset )
    }
    fmt.Println("")
}


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
	randword, _ := Getword()
	randword = strings.ToUpper(randword)
	// fmt.Println("The random word is: ", randword)
	fmt.Printf(green + "\n            WELCOME TO DADSON'S WORDLE GAME             \n     GUESS THE 5-LETTER WORD AND WIN A CASH PRICE      \n\n" + reset)

	// menu :=
	// `Commands
	// 1. New Game
	// 2. Score
	// 3. Quit
	// `
	// fmt.Println(menu)

    keyboard := CreateKeyboard()
    ShowKeyboard(keyboard)
	var guess string
	guesses_left := 7

	for i := 1; i < 7; i++ {
		guesses_left--
		fmt.Printf("\n\n[%v Guesses left] Guess the word: \n", guesses_left)
		fmt.Scan(&guess)
		for len(guess) != 5 {
			fmt.Printf("You entered a %d letter-word. Enter a 5 letter-word\n", len(guess))
			fmt.Printf("[%v Guesses left] Guess the word: \n", guesses_left)
			fmt.Scan(&guess)
		}
		guess = strings.ToUpper(guess)
		for !reg.MatchString(guess){
		    fmt.Printf("Your word contains invalid characters\n")
			fmt.Printf("[%v Guesses left] Guess the word: \n", guesses_left)
			fmt.Scan(&guess)
		}
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
		ShowKeyboard(keyboard)
	}
	fmt.Println("The word is: ", randword)
}
