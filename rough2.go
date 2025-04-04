package main

import (
	"fmt"
)

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
// 		player = getPlayerNames("20 chars exceeded. Enter player name again")
// 	}
// 	return
// }

func InWord2(char rune, word string) bool {
	for _, character := range word {
		if char == character {
			return true
		}
	}
	return false
}

func processGreens(guess string, word string) [5]string {
	response := [5]string{}
	for index, char := range guess {
		if rune(word[index]) == char {
			response[index] = "Green"
		}
	}
}

func ProcessGuess2(guess string, word string) [5]string {
	response := [5]string{}

	for index, char := range guess {
		if rune(word[index]) == char {
			response[index] = "Green"
		} else if InWord2(char, word) {
			response[index] = "Yellow"
		} else {
			response[index] = "Red"
		}
	}
	return response
}
func main() {
	word := "guest"
	guess := "stars"
	response := ProcessGuess2(guess, word)
	fmt.Println(response)
}
