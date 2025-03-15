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
    ages["Kobina"] = 65

    fmt.Println(ages)


}