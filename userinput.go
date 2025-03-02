package main

import "fmt"

// Reading simple user inputs
func main() {
    var response string
    fmt.Scan(&response)
    fmt.Println("User typed: " + response)
}