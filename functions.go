package main

import (
"fmt"
)

// Functions with no parameters with no return
func greet() {
    fmt.Println("Hello world! ")
}

// Functions with parameters and a return case 1: Unnamed return but return type specified
func add(first int, second int) int {
    return (first + second)
}

// Functions with specified parameter names and return names as well as types
func mod(first int, second int) (remainder int) {
    remainder = first % second
    return
}

// Functions with multiple returns
func divide(first int, second int) (quotient int, remainder int) {
    remainder = mod(first, second)
    quotient = first/second
    return
}


func main() {
    greet()
    num1, num2 := 20, 3
    result := add(num1, num2)
    modulo := mod(num1, num2)
    quotient, remainder := divide(num1, num2)
    fmt.Printf("sum is %v \n", result)
    fmt.Printf("modulo is: %v \n", modulo)
    fmt.Printf("Quotient: %v and Remainder %v", quotient, remainder)
}