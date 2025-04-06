package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
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

// func addScore(score string, player_name string, scoreData bool) {
// 	header := []string{"Rank", "Player Name", "Score", "DateTime"}

// 	if add_header {
// 		hs_file, err := os.OpenFile("Highscore.csv", os.O_CREATE|os.O_WRONLY, 0644)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		defer hs_file.Close()

// 		writer := csv.NewWriter(hs_file)
// 		writer.Write(header)
// 		curr_time := time.Now()
// 		formatted_time := curr_time.Format("2006-01-02 15:04:05")
// 		writer.Write([]string{"1", player_name, score, formatted_time})
// 	} else {
// 		hs_file, err := os.Open("Highscore.csv")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		defer hs_file.Close()

// 		scoreData, scores := getScoreInfo()

// 		writer := csv.NewWriter(hs_file)
// 		curr_time := time.Now()
// 		formatted_time := curr_time.Format("2006-01-02 15:04:05")
// 		writer.Write([]string{"1", player_name, score, formatted_time})
// 	}
// }

// func ProcessGuess(guess string, word string) [5]string {
// 	response := [5]string{}
//
// 	for index, char := range guess {
// 		if rune(word[index]) == char {
// 			response[index] = "Green"
// 		} else if InWord(char, word) {
// 			response[index] = "Yellow"
// 		} else {
// 			response[index] = "Red"
// 		}
// 	}
// 	return response
// }

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

	// 	array := Get_Dict()
	// 	fmt.Printf("%v %v", array, len(array))
	//
	// 	indict := WordIn_Dict("pumma", array)
	// 	fmt.Println(indict)

	// 	file, err := os.OpenFile("Highscore.ddl", os.O_WRONLY|os.O_APPEND, 0644)
	// 	if err != nil {
	// 	    fmt.Printf("This error occured: %v", err)
	// 	}
	// 	defer file.Close()

	csv_file, err := os.OpenFile("Highscore.csv", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	//     defer csv_file.Close()

	writer := csv.NewWriter(csv_file)

	header := []string{"Position", "Player Name", "Score", "DateTime"}

	scores := [][]string{
		{"1", "Ato", "34", "2025-03-22 15:35:59"},
		{"2", "Kwame", "53", "2022-03-09 18:41:12"},
		{"3", "Kofi", "23", "2025-04-12 14:51:12"},
		{"4", "Yofi", "67", "2025-06-22 09:32:12"},
		{"5", "Kweku", "79", "2025-08-21 15:35:12"},
	}

	writer.Write(header)
	for _, row := range scores {
		write_err := writer.Write(row)
		if write_err != nil {
			log.Fatal(write_err)
		}
	}
	writer.Flush()

	fmt.Printf("writer type: %T, csv_file type: %T", writer, csv_file)

}
