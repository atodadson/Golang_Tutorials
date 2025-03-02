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
// func main() {
//     var cars = [4]string{"volvo", "bugatti", "benz", "bentley"}
//
//     for index, item := range cars {
//         fmt.Printf("Index: %d, item: %s\n", index ,item)
//     }
// }

// Example, simple for loop (while) that breaks upon user's command
// func main() {
//     endLoop := true
//     command := ""
//     for endLoop {
//         fmt.Println("Type in your command: ")
//         fmt.Scan(&command)
//         if command == "quit" {
//             endLoop = false
//         }
//     }
// }

// Crontrolling the loop with break
// func main() {
//     arr := []int{-1, 0 , 1, 2, 3, 4}
//     for i := 0; i < 10; i++ {
//         fmt.Printf("Loop number %d\n", (i+1) )
//         if arr[i] > 3 {
//             fmt.Println("Condition satisfied, breaking loop.....")
//             break
//         }
//     }
// }

// Using continue in loops
func main() {
    arr := []string{"apple", "banana", "mango", "stone", "orange"}
    fmt.Println("I like eating: ")
    for index, item := range arr {
        if index == 3 {
            continue
        }
        /* The for loops just jumps to the next and does not excecute this
        line when the above condition is met */
        fmt.Printf("%v. %s\n", (index+1), item)
    }
}