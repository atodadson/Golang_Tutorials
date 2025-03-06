package main

// import (
// "fmt"
// )

// Using panic and recover
func Divide(first float64, second float64) (answer float64) {
    if second == 0 {
        panic("Can't Divide by 0")
    }
    return first / second

}

func main() {
    Divide(20, 0)
}