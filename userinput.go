package main

import "fmt"

// Reading simple user inputs
// func main() {
//     var response string
//     fmt.Scan(&response)
//     fmt.Println("User typed: " + response)
// }

// Reading mutlitple user inputs
func main() {
    var firstname, lastname string
    fmt.Println("Enter your firstname and lastname")
    fmt.Scan(&firstname, &lastname)
    fmt.Printf("Hello Mr./Mrs. %s %s", firstname, lastname )
}