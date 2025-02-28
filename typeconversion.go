package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	my_string := "Hello"
// 	my_bool := true
// 	my_int := 34
// 	my_float := 23.23232

// 	fmt.Printf("my_string is %s, my_bool is %t, my_int is %d, my_float is %f", my_string, my_bool, my_int, my_float)
// 	fmt.Printf("my_string has type %T, my_bool has type %T, my_int has type %T, my_float has type %T", my_string, my_bool, my_int, my_float)
// }

// func add(first int, second int) int {
// 	return first + second
// }

// func main() {
// 	no1, error1 := strconv.Atoi(os.Args[1])
// 	no2, error2 := strconv.Atoi(os.Args[2])
// 	var sum = add(no1, no2)

// 	fmt.Println(error1, error2)
// 	fmt.Println(sum)
// }

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Please provide a number as an argument!")
// 		return
// 	}

// 	no, err := strconv.Atoi(os.Args[1])
// 	fmt.Println(no)
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println("Couldn't convert: " + os.Args[1])
// 	} else {
// 		fmt.Println(no)
// 	}
// }

import (
	"os"
	"strconv"
)

// String to Integer conversion
// func main() {
// 	var no int = 100
// 	fmt.Println(reflect.TypeOf(no))

// 	var intStr string = "100"
// 	fourBaseEightBitInt, _ := strconv.ParseInt(intStr, 4, 64) // takes the input as base 4 and converts to in

// 	tenBaseSixteenBitInt, _ := strconv.ParseInt(intStr, 10, 16) // no 100, and

// 	fmt.Println(fourBaseEightBitInt)
// 	fmt.Println(tenBaseSixteenBitInt)
// 	fmt.Println(reflect.TypeOf(fourBaseEightBitInt))
// 	fmt.Println(reflect.TypeOf(tenBaseSixteenBitInt))
// }

// Integer to String conversion
// func main() {
// 	var noOfPlayers = 8
// 	strin := strconv.Itoa(noOfPlayers)
// 	fmt.Println(strin)
// }

//                             Challenge
// What happens if you run the program like so?
// go run typeconversion.go one two
// Handle any conversion error and call panic() if there's a conversion error.

func add(first int, second int) int {
	return first + second
}

func main() {
	if len(os.Args) != 2 {
		panic("Expected 2 command line arguments")
	}
	no1, err1 = strconv.Atoi(os.Args)
}
