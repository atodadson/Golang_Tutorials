package main

import "fmt"
/*
Assignment - build a phone book
Here are your contacts:
1 Alice 555-123
2 Bob 555-124
3 Jean 555-125
1 Welcome to your phonebook.
2 Command> store
3 Enter contact: Rob 555-126
4 Contact saved
5 Command> list
6 Alice 555-123
7 Bob 555-124
8 Jean 555-125
9 Rob 555-126
10 Command> lookup
11 Enter name: Alice
12 Alice has number: 555-123
HINT: you might need to use both a map and a slice.
*/

// Assignment solution
func createPhonebook() {
    var command string
    contacts := make(map[string]string)
    fmt.Println("Welcome to your phonebook")

    for {
        fmt.Print("Command> ")
        fmt.Scan(&command)
        if command == "store" {
            fmt.Print("Enter contact: ")
            var contact string
            var no string
            fmt.Scan(&contact, &no)
            _ , exists := contacts[contact]
            if !exists {
                contacts[contact] = no
                fmt.Println("Contact saved")
            } else {
                fmt.Println("Contact exists, overwrite? [Y/N]")
                var overwrite string
                fmt.Scan(&overwrite)
                if overwrite == "Y" {
                    contacts[contact] = no
                    fmt.Println("Contact saved")
                }
            }
        } else if command == "list" {
            for key, value := range contacts {
                fmt.Println(key, value)
            }
        } else if command == "lookup" {
            fmt.Print("Enter name: ")
            var contact string
            fmt.Scan(&contact)
            _, exists := contacts[contact]
            if exists {
                fmt.Println(contacts[contact])
            } else {
                fmt.Println("Contact doesn't exist. Do you want to save it [Y/N]")
                var save string
                fmt.Scan(&save)
                if save == "Y" {
                    fmt.Println("Enter contact number")
                    var no string
                    fmt.Scan(&no)
                    contacts[contact] = no
                } else {
                    continue
                }
            }
        } else if command == "quit" {
            break
        } else {
            fmt.Println("Unknown command: ", command)
        }
    }
    fmt.Println("Bye")
}


func main() {
//     /* Creating maps in go
//     To create a map in Go, we need to use the following syntax:
//     map[<key type>]<value type>{ ... entries  */
//     phonebook := map[int]string{1: "Kwame", 2:"Bio"}
//
//     fmt.Println(phonebook)
//
//     // We can create a map with make and later add entries
//     ages := make(map[string]int)
//
//     // Adding entries
//     ages["kofi"] = 34
//     ages["Kwame"] = 12
//     ages["Kwabena"] = 65
//
//     fmt.Println(ages)
//
//     // Checking if a key exists
//     key := "Kwame"
//     _ , exists := ages[key]
//     fmt.Printf("Key exists = %v \n", exists)
//
//     // You can use the above construct in an if statement - To correct
// //     if _, exists := ages["Kwabena"] {
// //         age = ages[key]
// //         fmt.Printf("His age is: %v", age)
// //     }
//
//     // Iterating over maps
//     for key, value := range ages {
//         fmt.Printf("%s's age is: %d", key, value)
//     }
//
//     // Deleting an entry
//     delete(ages, "Kofi")

    // Solution implementation
    createPhonebook()

}