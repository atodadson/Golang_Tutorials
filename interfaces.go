package main

import "fmt"

func main() {
    // While struct defines an object, interfaces defines behaviour
    // Declaration of interfaces
    type Point struct{
        x float32
        y float32
    }

    type Shape interface{
        area() float32
        location() Point
    }

}