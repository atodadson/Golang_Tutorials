package main

import "fmt"

// Reading simple user inputs
// func main() {
//     var response string
//     fmt.Scan(&response)
//     fmt.Println("User typed: " + response)
// }

// Reading mutlitple user inputs
// func main() {
//     var firstname, lastname string
//     fmt.Println("Enter your firstname and lastname")
//     fmt.Scan(&firstname, &lastname)
//     fmt.Printf("Firstname: %s Lastname: %s", firstname, lastname )
// }

/*
Try modifying the program so it takes several players first and then ensures all players get a name.
Here’s an example output of such a program:

1 Enter number of players:
2 3
3 Enter Player 1 name:
4 Alice
5 Enter Player 2 name:
6 Bob
7 Enter Player 3 name:
8 Jean
9 Players are: Alice Bob Jean
*/

func main() {
    fmt.Println("Enter number of players")
    var num_of_players int
    fmt.Scanf("%d", &num_of_players)
    var names string
    for i := 1; i <= num_of_players; i++ {
        var name string
        fmt.Printf("Enter Player %d name: \n", i)
        fmt.Scanf("%s", &name)
        names = names + " " + name
    }
    fmt.Printf("Names of players are: %v", names)
}