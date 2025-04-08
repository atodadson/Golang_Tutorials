package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
1. Speed race. Set time, get the most words
2. Slow but sure
*/

// ANSI escape codes for different colors
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var reset = "\033[0m"

// Keyboard display
func CreateKeyboard() (keyboard map[string]string) {
	keyboard = make(map[string]string, 6)
	for i := 'A'; i <= 'Z'; i++ {
		keyboard[string(i)] = reset
	}
	return keyboard
}

// Score values
var score_list = map[int]int{1: 13, 2: 8, 3: 5, 4: 3, 5: 2, 6: 1}

// Prints the current colouring of the keyboard
func ShowKeyboard(keyboard map[string]string) {
	row1 := []string{"       Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P"}
	row2 := []string{"       A", "S", "D", "F", "G", "H", "J", "K", "L"}
	row3 := []string{"        Z", "X", "C", "V", "B", "N", "M"}
	fmt.Println("")
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
	fmt.Printf("\n\n")
}

// Updates the colour of the keyboard after processing response
func UpdateKeyboard(letter string, colour string, keyboard map[string]string) map[string]string {
	if colour == "Green" {
		keyboard[letter] = green
	} else if colour == "Yellow" {
		if keyboard[letter] == reset {
			keyboard[letter] = yellow
		}
	} else if colour == "Red" {
		keyboard[letter] = red
	}
	return keyboard
}

// Returns all the words in the 5-letter words dictionary
func GetDict() (dictionary []string) {
	// Open the first file
	file, err := os.Open("dictionary5.txt")
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
	// fmt.Println(dictionary)
	return
}

// Returns a random word for the game words
func GetWord() (randword string) {
	// Open the first file
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
	randomIndex := rand.Intn(len(words))
	randword = words[randomIndex]
	return
}

// Checks if a letter is in a word
func InWord(char rune, word string) bool {
	for _, character := range word {
		if char == character {
			return true
		}
	}
	return false
}

// Checks if processed response checks green for all letters of the word
func allValuesGreen(m [5]string) bool {
	for _, value := range m {
		if value != "Green" {
			return false
		}
	}
	return true
}

// Checks if word is in file containing wordlist
func WordInDict(guess string, dictionary []string) bool {
	for _, word := range dictionary {
		if strings.ToUpper(guess) == strings.ToUpper(word) {
			return true
		}
	}
	return false
}

func numberLetters(word string) map[string][]int {
	var numbering = make(map[string][]int)
	for index, letter_rune := range word {
		letter := string(letter_rune)
		if _, exists := numbering[letter]; exists {
			//Just don't do anything
		} else {
			numbering[letter] = []int{(index + 1)}
		}
		for i := index + 1; i < len(word); i++ {
			check_rune := word[i]
			check_letter := string(check_rune)
			if check_letter == letter {
				numbering[letter] = append(numbering[letter], (i + 1))
			}
		}
	}
	return numbering
}

func removeElement(slice []int, item int) []int {
	index := -1
	for i, v := range slice {
		if v == item {
			index = i
			break
		}
	}
	if index == -1 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// ProcessGuess processes the guess and returns a response
func ProcessGuess(guess string, word string) [5]string {
	response := [5]string{}
	word_map := numberLetters(word)

	for index, letter_rune := range guess {
		letter := string(letter_rune)
		positions, exists := word_map[letter]
		if !exists {
			response[index] = "Red"
		} else {
			response[index] = "Yellow"
			for _, position := range positions {
				if position == (index + 1) {
					response[index] = "Green"
					if len(positions) == 1 {
						delete(word_map, letter)
					} else {
						word_map[letter] = removeElement(positions, (index + 1))
					}
				}
			}
		}
	}

	for index, letter_rune := range guess {
		letter := string(letter_rune)
		positions, exists := word_map[letter]
		if response[index] == "Green" {
			// Do nothng
		} else {
			if !exists {
				response[index] = "Red"
			} else {
				response[index] = "Yellow"
				if len(positions) == 1 {
					delete(word_map, letter)
				} else {
					word_map[letter] = removeElement(positions, (index + 1))
				}
			}
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

func showGameRules() {
	fmt.Println("\n<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Printf("                            HOW TO PLAY\n")
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	var GameRules = fmt.Sprintf(
		"1. Each guess must be a valid 5-letter word\n" +
			"2. The color of each letter will change to show how close your guess is to the answer\n\nExample: " +
			red + "W O R " + yellow + "D " + green + "S" + reset + " means:\n" +
			red + "'W', 'O', 'R'" + reset + " are not in the word\n" +
			yellow + "D" + reset + " is in the word but in the wrong spot\n" +
			green + "S" + reset + " is in the word and in the right spot\n")
	fmt.Println(GameRules)
}

func fileExists(filename string) bool {
	// Use os.Stat to get the file information
	_, err := os.Stat(filename)

	// Check if the error is of type "file does not exist"
	if os.IsNotExist(err) {
		return false
	}

	// If no error, the file exists
	return true
}

func SortData(data [][]string) [][]string {
	// Custom sorting function based on the age (index 2)
	sort.SliceStable(data, func(i, j int) bool {
		// Convert the age from string to integer for proper comparison
		ageI, errI := strconv.Atoi(data[i][1])
		ageJ, errJ := strconv.Atoi(data[j][1])

		// If there's an error in conversion, consider them equal for now
		if errI != nil || errJ != nil {
			return false
		}

		// Sort in descending order based on age
		return ageI > ageJ
	})

	// Update index 0 (ranking) based on the sorted order
	for i := range data {
		data[i][0] = strconv.Itoa(i + 1) // Update ranking with 1-based index
	}

	return data
}

// func getPlayerNames(message string) (player string) {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Println(message)
// 	// var container string
// 	// fmt.Scanf("%s", &container)
// 	player_name, err := reader.ReadString('\n')
// 	player = strings.TrimSpace(player_name)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for len(player) > 20 {
// 		player = getPlayerName("20 chars exceeded. Enter player name again")
// 	}
// 	return
// }

func isHighScore(newscore int, scores []int) bool {
	if newscore == 0 {
		return false
	} else {
		if len(scores) <= 5 {
			return true
		} else {
			for _, score := range scores {
				if newscore > score {
					return true
				}
			}
		}
		return false
	}
}

func highScorePosition(newscore int, scores []int) (pos int) {
	sort.Ints(scores)
	pos = 1
	for index, score := range scores {
		if score > newscore {
			pos = len(scores) - index + 1
			return
		}
	}
	return
}

func getScoreInfo() (scoreData [][]string, scores []int) {
	csv_file, err := os.Open("Highscore.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csv_file.Close()

	reader := csv.NewReader(csv_file)
	records, err := reader.ReadAll()
	for _, record := range records {
		// fmt.Printf("THis is record: %v", record)
		scoreData = append(scoreData, record)
		score, _ := strconv.Atoi(record[1])
		scores = append(scores, score)
	}
	scoreData, scores = scoreData[1:], scores[1:]
	return
}

func addFirstScore(score string) {
	header := []string{"Rank", "Score", "DateTime"}
	hs_file, err := os.OpenFile("Highscore.csv", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}
	defer hs_file.Close()

	writer := csv.NewWriter(hs_file)
	writer.Write(header)
	curr_time := time.Now()
	formatted_time := curr_time.Format("2006-01-02 15:04")
	writer.Write([]string{"1", score, formatted_time})
	writer.Flush()
}

func addScores(scoreData [][]string) {
	header := []string{"Rank", "Score", "DateTime"}
	hs_file, err := os.OpenFile("Highscore.csv", os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}
	defer hs_file.Close()

	writer := csv.NewWriter(hs_file)
	writer.Write(header)
	for index, record := range scoreData {
		if index > 4 {
			break
		}
		writer.Write(record)
	}
	writer.Flush()
}

func updateHighScore(scorevalue int) {
	rankings := map[int]string{
		1: "1st",
		2: "2nd",
		3: "3rd",
		4: "4th",
		5: "5th",
	}
	score := strconv.Itoa(scorevalue)

	if !fileExists("Highscore.csv") {
		addFirstScore(score)
		fmt.Printf("You got a new High Score. You placed 1st\n     <<<< Your Score: %v >>>>\n", score)

	} else {
		scoreData, scores := getScoreInfo()
		if isHighScore(scorevalue, scores) {
			curr_time := time.Now()
			formatted_time := curr_time.Format("2006-01-02 15:04:05")
			new_record := []string{"1", score, formatted_time}
			scoreData = append(scoreData, new_record)
			scoreData = SortData(scoreData)

			addScores(scoreData)
			pos := highScorePosition(scorevalue, scores)
			rank := rankings[pos]
			fmt.Printf("You got a new High Score. You placed %s\n     <<<< Your Score: %v >>>>\n", rank, score)
		} else {
			fmt.Printf("     <<<< Your Score: %v >>>>\n", score)
		}
	}

}

func showHighScores() {
	filename := "Highscore.csv"
	fmt.Println("\n<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Printf("                             HIGH SCORES\n")
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	if !fileExists(filename) {
		fmt.Println("No High Score recorded yet")
	} else {
		scoreData, _ := getScoreInfo()
		for i := 0; i < len(scoreData); i++ {
			record := scoreData[i]
			fmt.Printf("\t %v. %v\t %v\n", record[0], record[1], record[2])
		}
	}
}

func showMenu() {
	fmt.Println("\n=====================================================================")
	fmt.Printf("                             MENU\n")
	fmt.Println("=====================================================================")
	menu :=
		`
1. Play
2. How to Play
3. High Scores
4. Quit

Enter 1,2,3 or 4`
	fmt.Println(menu)
}

func main() {
	// fmt.Println("The random word is: ", randword)
	fmt.Printf(green + "\n            WELCOME TO DADSON'S WORDLE GAME\n" + reset)

	var command string
	var playAgain = "y"

	showMenu()
	fmt.Scanf("%s\n", &command)
	for true {
		var guessed_right = true
		if command == "4" || strings.ToUpper(command) == "QUIT" {
			break
		} else if command == "2" {
			showGameRules()
			showMenu()
			fmt.Scanf("%s\n", &command)
		} else if command == "3" {
			showHighScores()
			showMenu()
			fmt.Scanf("%s\n", &command)
		} else if command == "1" {
			if strings.ToLower(playAgain) != "y" {
				break
			}

			level := 0
			var randword string
			var dictionary = GetDict()
			var score = 0

			for true {
				if !guessed_right {
					// Print out the result when game ends
					fmt.Printf("\nSorry, you lost!\n")
					fmt.Printf("\nThe word is: %v \n", randword)
					updateHighScore(score)
					fmt.Println("Do you want to play again [y/n]")
					fmt.Scanf("\n%s\n", &playAgain)
					break
				}
				level++
				fmt.Println("\n_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_")
				fmt.Printf("                            WORD %v\n", level)
				fmt.Printf("                      CURRENT SCORE: %v\n", score)
				fmt.Println("_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_")

				// Declaration of variable for each level
				keyboard := CreateKeyboard()
				// ShowKeyboard(keyboard)
				var guess string
				guesses_left := 7
				randword = GetWord()
				randword = strings.ToUpper(randword)
				// fmt.Println("The word is: ", randword)

				for i := 1; i < 7; i++ {
					guesses_left--
					fmt.Printf("\n[Guess %v/6] Enter word: \n", i)
					fmt.Scan(&guess)

					// Check For correct input
					for true {
						if len(guess) != 5 { // Word must be five letters
							fmt.Printf("You entered a %d letter-word. Enter a 5 letter-word\n", len(guess))
						} else if !WordInDict(guess, dictionary) { // Word must be in the dictionary
							fmt.Printf("Word does not exist. Try again\n")
						}
						if (len(guess) == 5) && (WordInDict(guess, dictionary)) {
							break
						} else {
							fmt.Printf("\n[Guess %v/6] Enter word: \n", i)
							fmt.Scan(&guess)
						}

					}

					guess = strings.ToUpper(guess)
					response := ProcessGuess(guess, randword)
					for i := 0; i < 5; i++ {
						keyboard = UpdateKeyboard(string(guess[i]), response[i], keyboard)
					}
					ShowKeyboard(keyboard)

					fmt.Printf(GetColour(response[0]) + string(guess[0]) + " " + reset)
					fmt.Printf(GetColour(response[1]) + string(guess[1]) + " " + reset)
					fmt.Printf(GetColour(response[2]) + string(guess[2]) + " " + reset)
					fmt.Printf(GetColour(response[3]) + string(guess[3]) + " " + reset)
					fmt.Printf(GetColour(response[4]) + string(guess[4]) + " " + "\n" + reset)
					if allValuesGreen(response) {
						fmt.Println("Yeeey, You got it right. That's a great guess. Progress to the next word")
						guessed_right = true
						score = score + score_list[i]
						break
					}
					guessed_right = false
				}
			}
			if strings.ToLower(playAgain) == "y" {
				// Do not show menu again
			} else {
				showMenu()
				fmt.Scanf("%s\n", &command)
				playAgain = "y"
			}
		} else {
			fmt.Println("Unexpected command. Enter One of (1,2,3 or 4)")
		}
	}
}
