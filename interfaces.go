package main

import "fmt"

// Implementation of interfaces
type Rectangle struct {
    length float32
    width float32
}

func (r Rectangle) area() float32 {
    return r.length * r.width
}

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

   myRect := Rectangle{20, 15}
   fmt.Printf("Area of rectangle with sides %v and %v is: %v", myRect.length, myRect.width, myRect.area())

}