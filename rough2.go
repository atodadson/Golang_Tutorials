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

package main

import (
	"fmt"
	"os"
	"runtime"
	// "encoding/json"
)

type HighScore struct {
	Rank     int `json:"rank"`
	Score    int `json:"score"`
	DateTime int `json:"datetime"`
}

func getHighScorePath() (filepath string) {
	var base string
	if runtime.GOOS == "windows" {
		base = os.Getenv("APPDATA")
	} else {
		base = os.Getenv("HOME")
	}
	filepath = base
	return
}

func main() {
	filepath := getHighScorePath()
	fmt.Println(filepath)
	fmt.Printf()
}
