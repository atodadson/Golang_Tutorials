package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    var player string
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Enter all names")
    player, err := reader.ReadString('\n')
    player = strings.TrimSpace(player)
    if err != nil {
        fmt.Println(err)
    }
//     fmt.Println("Enter player name")
//     fmt.ReadString(&player)
    fmt.Println(player)
}