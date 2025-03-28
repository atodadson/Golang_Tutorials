package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime/debug"
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

// Defining errors
var NumTooSmall = errors.New("Expected a number greater than zero")

func ReturnPositive(no int) (int, error) {
	if no > 0 {
		return no, nil
	} else {
		return 0, NumTooSmall
	}
}

func main() {
	Divide(20, 0)
	num, err := ReturnPositive(-12)
	check := debug.Stack()
	fmt.Println(check)

	// Error logging
	f, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(f)

	fmt.Printf("Num: %v, Error: %v", num, err)
}
