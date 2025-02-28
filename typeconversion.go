package main

import (
	"fmt"
	"os"
	"strconv"
)

// func main() {
// 	my_string := "Hello"
// 	my_bool := true
// 	my_int := 34
// 	my_float := 23.23232

// 	fmt.Printf("my_string is %s, my_bool is %t, my_int is %d, my_float is %f", my_string, my_bool, my_int, my_float)
// 	fmt.Printf("my_string has type %T, my_bool has type %T, my_int has type %T, my_float has type %T", my_string, my_bool, my_int, my_float)
// }

func add(first int, second int) int {
	return first + second
}

func main() {
	no1, error1 := strconv.Atoi(os.Args[1])
	no2, error2 := strconv.Atoi(os.Args[2])
	var sum = add(no1, no2)

	fmt.Println(error1, error2)
	fmt.Println(sum)
}
