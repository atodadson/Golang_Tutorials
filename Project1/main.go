package main

import (
    "fmt"
    "go_project1/helper"
    "github.com/softchris/math" // Importing an external module
)

func main() {
    helper.Help()
    helper.Greet()
    fmt.Println(math.Add(23,45))
}