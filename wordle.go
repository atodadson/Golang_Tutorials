package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
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
		fmt.Printf(keyboard[letter] + letter + " " + reset)
	}
	fmt.Println("")
	for _, letter := range row2 {
		fmt.Printf(keyboard[letter] + letter + " " + reset)
	}
	fmt.Println("")
	for _, letter := range row3 {
		fmt.Printf(keyboard[letter] + letter + " " + reset)
	}
}

func UpdateKeyboard(letter string, colour string, keyboard map[string]string) {
	if colour == "Green" {
		keyboard[letter] = green
	} else if colour == "Yellow" {
		if keyboard[letter] == reset {
			keyboard[letter] = yellow
		}
	} else if colour == "Red" {
		keyboard[letter] = red
	}
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

var GameRules = fmt.Sprintf("1. Each guess must be a valid 5-letter word\nExample: " +
	"2. The color of each letter will change to show how close your guess was to the word\n" +
	red + "W O R " + yellow + "D " + green + "S" + reset + "means:\n" +
	"'W', 'O' and 'R' are not in the word\n" +
	"D is in the word but in the wrong spot\n" +
	"S is in the word and in the right spot")

func main() {
	randword, _ := Getword()
	randword = strings.ToUpper(randword)
	// fmt.Println("The random word is: ", randword)
	fmt.Printf(green + "\n            WELCOME" + yellow + " TO" + red + " DADSON'S" + yellow + " WORDLE" + green + " GAME\n" + reset)

	menu :=
		`Commands
	1. Play
	2. How to Play
	3. High Scores
	4. Quit

	Enter 1,2,3 or 4
	`
	var command string

	for true {
		fmt.Println(menu)
		fmt.Scanf("%s", &command)
		if command == "4" || strings.ToUpper(command) == "QUIT" {
			break
		} else if command == "2" {
			fmt.Println(GameRules)
		} else if command == "3" {
			fmt.Println("Not up yet")
		} else if command == "1" {
			for true {
				keyboard := CreateKeyboard()
				ShowKeyboard(keyboard)
				var guess string
				guesses_left := 7

				var guessed_right = true

				if !guessed_right {
					break
				}

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
					for !reg.MatchString(guess) {
						fmt.Printf("Your word contains invalid characters\n")
						fmt.Printf("[%v Guesses left] Guess the word: \n", guesses_left)
						fmt.Scan(&guess)
					}
					response := ProcessGuess(guess, randword)
					for i := 0; i < 5; i++ {
						UpdateKeyboard(string(guess[i]), response[i], keyboard)
					}

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
			}

		} else {
			fmt.Println("Unexpected command. Enter One of (1,2,3 or 4)")
		}
		// fmt.Println("The word is: ", randword)
	}
}
