package main

import "fmt"

func main() {
    myArray := []string{"pineapple", "orange", "mango", "pawpaw", "sausage", "beef"}
    capacity := cap(myArray)
    length := len(myArray)
    fruits := myArray[1:6]

    fmt.Printf("Array: %v, Length: %v, Cap is %v, Slice: %v",myArray, length, capacity, fruits)

}