package main

import "fmt"

// 1. Basic For loops. Starting point and increment upto a limit
// func main() {
//     for i := 0; i < 10; i++ {
//         fmt.Println(i)
//     }
// }

// 2. The Go language does not a keyword for while. Instead it uses a modified for loop
// func main() {
//     i := 0
//     for i < 10 {
//         i = i + 1
//         fmt.Printf("The current value of i is: %d\n", i)
//     }
// }

// 3. The for each structure
func main() {
    var cars = [4]string{"volvo", "bugatti", "benz", "bentley"}

    for index, item := range cars {
        fmt.Printf("Index: %d, item: %s\n", index ,item)
    }

}