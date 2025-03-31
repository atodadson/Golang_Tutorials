package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func getPlayerName(message string) (player string) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println(message)
    player, err := reader.ReadString('\n')
    player = strings.TrimSpace(player)
    if err != nil {
        fmt.Println(err)
    }
    return
}

func main() {
    player := getPlayerName("I know your name: ")
    fmt.Println(player)
}