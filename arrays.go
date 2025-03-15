package main

import "fmt"

func main() {
    myArray := []string{"pineapple", "orange", "audi", "bugatti", "mango", "pawpaw", "mutton", "beef", "chicken"}
    capacity := cap(myArray)
    length := len(myArray)
//     fruits := myArray[1:6]

    // Removing element. case1. Creating a new array from slice
    cars_start, cars_end := 2, 4
    cars := myArray[cars_start:cars_end]

    fruits_start1, fruits_end1, fruits_start2, fruits_end2 := 0, 2, 4, 6
    fruits := append(myArray[fruits_start1:fruits_end1], myArray[fruits_start2:fruits_end2]...)//ellipsis to unpack

    fmt.Printf("Array: %v, Length: %v, Cap is %v \n" ,myArray, length, capacity)
    fmt.Printf("Array: %v, Cars: %v, Fruits is %v \n" ,myArray, cars, fruits)

}