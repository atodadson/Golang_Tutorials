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

// func processGreens(guess string, word string) [5]string {
// 	response := [5]string{}
// 	for index, char := range guess {
// 		if rune(word[index]) == char {
// 			response[index] = "Green"
// 		}
// 	}
// }

func numberLetters (word string) (map[string][]int){
    var numbering = make(map[string][]int)
    for index, letter_rune := range word {
        letter := string(letter_rune)
        if _, exists := numbering[letter]; exists {
            //Just don't do anything
            // numbering[letter] = append(numbering[letter], (index+1) )
        } else {
            numbering[letter] = []int{(index + 1)}
        }
        for i := index+ 1; i < len(word); i++ {
            check_rune := word[i]
            check_letter := string(check_rune)
            if check_letter == letter {
                numbering[letter] = append(numbering[letter], (i+1))
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
    fmt.Printf("item: %v deleted", item)
    return append(slice[:index], slice[index+1:]...)
}

func ProcessGuess2(guess string, word string) [5]string {
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
	            fmt.Printf("position: %T, index: %T\n", position, (index+1))
	            if position == (index + 1) {
	                response[index] = "Green"
	                if len(positions) == 1 {
	                    delete(word_map, letter)
	                } else {
	                    word_map[letter] = removeElement(positions, (index + 1) )
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
                    word_map[letter] = removeElement(positions, (index + 1) )
                }
            }
	    }

	}
	return response
}

func main() {
	word := "guess"
	guess := "gssss"
	response := ProcessGuess2(guess, word)
	fmt.Println(response)
	fmt.Println(word[2])
	mymap := make(map[int]int)
	mymap[34] = 43
	mymap[12] = 21
	check, err := mymap[23]
	fmt.Println(check)
	fmt.Println(err)
// 	numberLetters(word)
	result := numberLetters(word)
	fmt.Println(result)
	answer := ProcessGuess2(guess, word)
	fmt.Println(answer)
}
