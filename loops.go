package main

import "fmt"

// Basic For loops. Starting point and increment upto a limit
// func main() {
//     for i := 0; i < 10; i++ {
//         fmt.Println(i)
//     }
// }

// The Go language does not a keyword for while. Instead it uses a modified for loop
func main() {
    i := 0
    for i < 10 {
        i = i + 1
        fmt.Printf("The current value of i is: %d\n", i)
    }
}
