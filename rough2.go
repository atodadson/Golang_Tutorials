package main

import (
	"fmt"
	"encoding/json"
	"os"
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

func main() {
    filename := "products.json"
    json_file, _ := os.Open(filename)
    json_map, _ := json.Unmarshal(json_file)
    fmt.Printf("%v", json_map)

}