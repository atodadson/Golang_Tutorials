package main

import "fmt"

// Implementation of interfaces
type Rectangle struct {
    length float32
    width float32
}

type Point struct{
    x float32
    y float32
}

func (r Rectangle) area() float32 {
    return r.length * r.width
}

func (r Rectangle) location() Point {
    return Point{r.length, r.width}
}

func main() {
    // While struct defines an object, interfaces defines behaviour
    // Declaration of interfaces

    type Shape interface{
        area() float32
        location() Point
    }

   myRect := Rectangle{20, 15}
   fmt.Printf("Area of rectangle with sides %v and %v is: %v \n", myRect.length, myRect.width, myRect.area())
   fmt.Printf("Rectangle with sides %v and %v has location: %v \n", myRect.length, myRect.width, myRect.location())

}