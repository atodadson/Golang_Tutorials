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

var scoreData = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

func main() {
	for _, record := range scoreData {
		fmt.Printf("\t %v\t %v\t %v\n", record[0], record[1], record[2])
	}
}
