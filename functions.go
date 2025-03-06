package main

import "fmt"

// Functions with no parameters with no return
func greet() {
    fmt.Println("Hello world! ")
}

// Functions with parameters and a return case 1: Unnamed return but return type specified
func add(first int, second int) int {
    return (first + second)
}


func mod(first int, second int) (remainder int) {
    remainder = first % second
    return
}

func main() {
    greet()
    result := add(20, 5)
    fmt.Printf("result is %v", result)
}