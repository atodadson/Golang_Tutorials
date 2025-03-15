// package main
//
// import "fmt"
//
// // Assignment solution
// func logger() {
//     input_list := make([]string, 0)
//     for true {
//         fmt.Println("Command>")
//         var command string
//         var user_input string
//         fmt.Scanf("%s\n", &command)
//         if command == "quit" {
//             break
//         } else if command == "new"{
//             fmt.Printf("Enter prompt: ")
//             fmt.Scanf("%s\n", &user_input)
//             input_list = append(input_list, []string{user_input}...)
//         } else if command == "list" {
//             for index, item := range input_list {
//                 fmt.Printf("%d. %s \n", index, item)
//             }
//         } else {
//             fmt.Println("You entered the wrong command")
//         }
//
//     }
//
//
// }
//
// func main() {
// //     myArray := []string{"pineapple", "orange", "audi", "bugatti", "mango", "pawpaw", "mutton", "beef", "chicken"}
// //     capacity := cap(myArray)
// //     length := len(myArray)
// // //     fruits := myArray[1:6]
// //
// //     // Removing element. case1. Creating a new array from slice
// //     cars_start, cars_end := 2, 4
// //     cars := myArray[cars_start:cars_end]
// //
// //     fruits_start1, fruits_end1, fruits_start2, fruits_end2 := 0, 2, 4, 6
// //     fruits := append(myArray[fruits_start1:fruits_end1], myArray[fruits_start2:fruits_end2]...)//ellipsis to unpack
// //
// //     fmt.Printf("Array: %v, Length: %v, Cap is %v \n" ,myArray, length, capacity)
// //     fmt.Printf("Array: %v, Cars: %v, Fruits is %v \n" ,myArray, cars, fruits)
// //
// //     // Creating Slices with make()
// //     slice := make([]int, 3, 5) // Create array make(type, length, capacity)
// //     fmt.Printf("Slice: %v ", slice)
// //
// //     // Making copies
// //     fruits_copy := make([]string, 8)
// //     copy(fruits_copy, fruits[0:3])
// //     fmt.Printf("Fruits_copy: %v ", fruits_copy)
//
//     logger()
//
// }
//
// /*
// Assignment - store log entries
// Create an array meant for log entries. It can be used in the following way:
// 1 command> new
// 2 here's a new entry
// 3 command> new
// 4 here's another entry
// 5 command> list
// 6 here's a new entry
// 7 here's another entry
// 8 command> quit
// 9 bye
// */


// Solution from book
package main

import ("fmt")

func main() {
    // create array
    arr := make([]string, 0)
    var response string
    for {
        fmt.Print("command> ")
        fmt.Scan(&response)
        if response == "quit" {
            break
        } else if response == "new" {
            fmt.Print("Entry:")
            fmt.Scan(&response)

            // save entry to list
            arr = append(arr, response)
            fmt.Println("Saving entry")
        } else if response == "list" {
            // list entries
            fmt.Println("Listing entries")
            for i := 0; i < len(arr); i++ {
            fmt.Println(arr[i])
            }
        } else {
            fmt.Println("Unknown command", response)
        }
    }
    fmt.Println("bye")
}
