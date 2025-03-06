package main

import (
"fmt"
)

// Using recover and defer
func ErrorHandler() {
    if r := recover(); r != nil {
        fmt.Printf("Recovering from error: %v \n", r)
    }
}

// Using panic and recover
func Divide(dividend float64, divisor float64) (answer float64) {
    defer ErrorHandler()
    if divisor == 0 {
        panic("Can't Divide by 0")
    }
    return dividend / divisor
}


func main() {
    Divide(20, 0)
}