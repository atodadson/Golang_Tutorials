package main

import "fmt"

func main() {
    /* Creating maps in go
    To create a map in Go, we need to use the following syntax:
    map[<key type>]<value type>{ ... entries  */
    phonebook := map[int]string{1: "Kwame", 2:"Bio"}

    fmt.Println(phonebook)

    // We can create a map with make and later add entries
    ages := make(map[string]int)

    // Adding entries
    ages["kofi"] = 34
    ages["Kwame"] = 12
    ages["Kwabena"] = 65

    fmt.Println(ages)

    // Checking if a key exists
    key := "Kwame"
    _ , exists := ages[key]
    fmt.Printf("Key exists = %v \n", exists)

    // You can use the above construct in an if statement - To correct
//     if _, exists := ages["Kwabena"] {
//         age = ages[key]
//         fmt.Printf("His age is: %v", age)
//     }

    // Iterating over maps
    for key, value := range ages {
        fmt.Printf("%s's age is: %d", key, value)
    }

}